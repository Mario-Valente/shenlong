package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateCronsK8s(name string, namespace string, image string, command []string, pathKubeconfig string, ttl int32, schedule string) (*batchv1.CronJob, error) {
	clientK8s, err := CreateClientK8s(pathKubeconfig)
	if err != nil {
		fmt.Println("error to create client in k8s")
	}

	cron := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batchv1.CronJobSpec{
			Schedule:          schedule,
			ConcurrencyPolicy: batchv1.ForbidConcurrent,
			JobTemplate: batchv1.JobTemplateSpec{
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
			},
		},
	}

	cronClient := clientK8s.BatchV1().CronJobs(namespace)
	result, err := cronClient.Create(context.TODO(), cron, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating Cron: %s", err.Error())
	}

	return result, nil
}

func GetCronsK8s(name string, namespace string, pathKubeconfig string) (*batchv1.CronJob, error) {
	clientK8s, err := CreateClientK8s(pathKubeconfig)
	if err != nil {
		fmt.Println("error to create client in k8s")
	}

	cronClient := clientK8s.BatchV1().CronJobs(namespace)
	cron, err := cronClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting Cron: %s", err.Error())
	}

	cronJSON, err := json.MarshalIndent(cron, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling Cron to JSON: %s", err.Error())
	}
	fmt.Printf("Cron: %s\n", string(cronJSON))

	return cron, nil
}

func DeleteCronsK8s(name string, namespace string, pathKubeconfig string) error {
	clientK8s, err := CreateClientK8s(pathKubeconfig)
	if err != nil {
		fmt.Println("error to create client in k8s")
	}

	cronClient := clientK8s.BatchV1().CronJobs(namespace)
	err = cronClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("Error deleting Cron: %s", err.Error())
	}

	return nil
}
