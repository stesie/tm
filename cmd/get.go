package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve resources from k8s cluster",
}

// TODO: single interface for resulting object

// servicesCmd represents the builds command
var listServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List of knative service resources",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listServices()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

// listBuildsCmd represents the builds command
var listBuildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "List of knative build resources",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listBuilds()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

// listRoutesCmd represents the builds command
var listRoutesCmd = &cobra.Command{
	Use:   "routes",
	Short: "List of knative routes resources",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listRoutes()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

var listRevisionsCmd = &cobra.Command{
	Use:   "revisions",
	Short: "List of knative revision resources",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listRevisions()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

var listPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "List of pods",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listPods()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

var listBuildTemplatesCmd = &cobra.Command{
	Use:   "buildtemplates",
	Short: "List of buildtemplates",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listBuildTemplates()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

var listConfigurationsCmd = &cobra.Command{
	Use:   "configurations",
	Short: "List of configurations",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := listConfigurations()
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(listServicesCmd)
	getCmd.AddCommand(listBuildsCmd)
	getCmd.AddCommand(listRoutesCmd)
	getCmd.AddCommand(listRevisionsCmd)
	getCmd.AddCommand(listPodsCmd)
	getCmd.AddCommand(listBuildTemplatesCmd)
	getCmd.AddCommand(listConfigurationsCmd)
}

func format(v interface{}) ([]byte, error) {
	switch output {
	case "json":
		return json.MarshalIndent(v, "", "    ")
	case "yaml":
		return yaml.Marshal(v)
	}
	return []byte{}, nil
}

func listBuilds() (string, error) {
	list, err := build.BuildV1alpha1().Builds(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listServices() (string, error) {
	list, err := serving.ServingV1alpha1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listRoutes() (string, error) {
	list, err := serving.ServingV1alpha1().Routes(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s%-25s", "NAME", "NAMESPACE", "TARGETS(%)")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s%v", str, item.Name, item.Namespace, item.Spec.Traffic)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listRevisions() (string, error) {
	list, err := serving.ServingV1alpha1().Revisions(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listPods() (string, error) {
	list, err := core.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listBuildTemplates() (string, error) {
	list, err := build.BuildV1alpha1().BuildTemplates(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}

func listConfigurations() (string, error) {
	list, err := serving.ServingV1alpha1().Configurations(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if output == "" {
		str := fmt.Sprintf("%-25s%-25s", "NAME", "NAMESPACE")
		for _, item := range list.Items {
			str = fmt.Sprintf("%s\n%-25s%-25s", str, item.Name, item.Namespace)
		}
		return str, err
	}
	o, err := format(list)
	return string(o), err
}
