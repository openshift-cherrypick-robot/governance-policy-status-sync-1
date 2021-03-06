// Copyright (c) 2020 Red Hat, Inc.

package e2e

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Test cluster ns creation", func() {
	It("Should contain label on cluster ns", func() {
		ns, err := clientManaged.CoreV1().Namespaces().Get(testNamespace, metav1.GetOptions{})
		Expect(err).To(BeNil())
		Expect(ns.GetLabels()["policy.open-cluster-management.io/isClusterNamespace"]).To(Equal("true"))
	})
})
