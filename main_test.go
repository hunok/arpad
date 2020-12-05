package main

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestImages(t *testing.T) {
	t.Parallel()

	fakeClient := fake.NewSimpleClientset()

	fakeClient.CoreV1().Pods("default").Create(context.TODO(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-test-pod",
			Namespace: "default",
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "busybox",
					Image:           "busybox:v0.1",
					ImagePullPolicy: v1.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
		},
	}, metav1.CreateOptions{})

	testData, err := listImages(fakeClient, "")

	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}

	if len(testData) != 1 {
		t.Errorf("Expected %d, got %d", 1, len(testData))
	}

}
