package k8s

import (
	"context"
	"fmt"
	"log"

	"github.com/Mario-valente/shenlong/pkgs/utilis"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateJobsK8s(name string, namespace string, image string, command []string, ttl int32, pathKubeconfig string) (string, error) {
	clientK8s, err := CreateClientK8s(pathKubeconfig)
	if err != nil {
		fmt.Println("error to create client in k8s")
	}

	name = utilis.CreateRadomName(name)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: &ttl,
			BackoffLimit:            new(int32),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyOnFailure,
					Containers: []corev1.Container{
						{
							Name:    name,
							Image:   image,
							Command: command,
						},
					},
				},
			},
		},
	}

	jobClient := clientK8s.BatchV1().Jobs(namespace)
	result, err := jobClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating Job: %s", err.Error())
	}
	fmt.Printf("Created Job %q.\n", result.GetObjectMeta().GetName())

	for {
		job, err := jobClient.Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			log.Fatalf("Error getting Job: %s", err.Error())
		}
		if job.Status.Succeeded > 0 {
			fmt.Println("Job completed")
			break
		}
	}

	return "Job created", nil

}

func GetJobsK8s(name string, namespace string, pathKubeconfig string) (string, error) {
	clientK8s, err := CreateClientK8s(pathKubeconfig)
	if err != nil {
		fmt.Println("error to create client in k8s")
	}

	jobClient := clientK8s.BatchV1().Jobs(namespace)
	job, err := jobClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting Job: %s", err.Error())
	}
	fmt.Printf("Job %q.\n", job.GetObjectMeta().GetName())

	return "Job retrieved", nil
}
