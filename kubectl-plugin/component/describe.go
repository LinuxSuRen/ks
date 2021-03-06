package component

import (
	"fmt"
	"github.com/linuxsuren/ks/kubectl-plugin/common"
	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func newComponentDescribeCmd(client dynamic.Interface, _ *kubernetes.Clientset) (cmd *cobra.Command) {
	opt := &describeOption{
		Option{
			Client: client,
		},
	}

	cmd = &cobra.Command{
		Use:               "describe",
		Short:             "Wrapper of kubectl describe",
		Aliases:           []string{"desc", "inspect"},
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: common.KubeSphereDeploymentCompletion(),
		PreRunE:           opt.preRunE,
		RunE:              opt.runE,
	}
	return
}

type describeOption struct {
	Option
}

func (o *describeOption) preRunE(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		o.Name = args[0]
	}

	if o.Name == "" {
		err = fmt.Errorf("please provide the name of component")
	}
	return
}

func (o *describeOption) runE(_ *cobra.Command, args []string) (err error) {
	var podName string
	var ns string
	if ns, podName, err = o.getPod(args[0]); err == nil {
		err = common.ExecCommand("kubectl", "describe", "-n", ns, "pod", podName)
	}
	return
}
