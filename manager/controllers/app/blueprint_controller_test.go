// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0
package app

import (
	"context"
	"time"

	app "github.com/mesh-for-data/mesh-for-data/manager/apis/app/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func deployBlueprint(namespace string, shouldSucceed bool) {
	const timeout = time.Second * 30
	const interval = time.Millisecond * 100

	blueprint := &app.Blueprint{}
	Expect(readObjectFromFile("../../testdata/blueprint-read.yaml", blueprint)).ToNot(HaveOccurred())

	// Set the correct namespace
	blueprint.SetNamespace(namespace)

	blueprintKey := client.ObjectKeyFromObject(blueprint)

	// Create Blueprint
	Expect(k8sClient.Create(context.Background(), blueprint)).Should(Succeed())

	// Ensure getting cleaned up after tests finish
	defer func() {
		blueprint := &app.Blueprint{ObjectMeta: metav1.ObjectMeta{Namespace: blueprintKey.Namespace, Name: blueprintKey.Name}}
		_ = k8sClient.Get(context.Background(), blueprintKey, blueprint)
		_ = k8sClient.Delete(context.Background(), blueprint)
	}()

	By("Expecting blueprint to be created")
	Eventually(func() error {
		return k8sClient.Get(context.Background(), blueprintKey, blueprint)
	}, timeout, interval).Should(Succeed())

	if shouldSucceed {
		By("Expecting Blueprint to eventually be ready")
		Eventually(func() bool {
			Expect(k8sClient.Get(context.Background(), blueprintKey, blueprint)).To(Succeed())
			return blueprint.Status.ObservedState.Ready
		}, timeout, interval).Should(BeTrue(), "Blueprint is not ready after timeout!")
	} else {
		By("Expecting Blueprint to never be ready because reconcile should not be called")
		Eventually(func() bool {
			Expect(k8sClient.Get(context.Background(), blueprintKey, blueprint)).To(Succeed())
			return blueprint.Status.ObservedState.Ready
		}, timeout, interval).Should(BeFalse(), "Blueprint should not be ready because reconciler should not have been invoked!")
	}
}

var _ = Describe("Blueprint Controller Real Env", func() {
	Context("Blueprint", func() {
		BeforeEach(func() {
			// Add any setup steps that needs to be executed before each test
		})

		AfterEach(func() {
			// Add any teardown steps that needs to be executed after each test
		})

		// Blueprints are successfully reconciled when deployed to m4d-blueprints only
		It("Test Blueprint Deploy to Correct Namespace", func() {
			deployBlueprint("m4d-blueprints", true)
		})

		// Blueprints not deployed to m4d-blueprints should not be successfully reconciled due to the filter preventing
		// reconcile from being called.
		It("Test Blueprint Deploy to Bad Namespace", func() {
			deployBlueprint("default", false)

		})
	})
})
