package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/saedx1/go-chainlets/chainlets"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	excludePkg []string
	direct     bool
	rootCmd    = &cobra.Command{
		Use:   "go-chainlets [FILE] [PACKAGE_NAME]",
		Short: "Returns all chains that end with a specific package",
		Long:  ``,
		Run:   runChainlets,
		Args:  cobra.MinimumNArgs(2),
	}
)

func runChainlets(cmd *cobra.Command, args []string) {
	filename := args[0]
	lookForPkg := args[1]

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	graphStr := string(bytes)

	graph := chainlets.StrToGraph(graphStr)
	graph.ExcludePkgs(excludePkg)

	circularDeps := graph.CircularDep()
	if len(circularDeps) != 0 {
		fmt.Println("Found the following circular deps")
		fmt.Println(circularDeps)
		os.Exit(1)
	}

	chains := graph.Chains(lookForPkg)

	if len(chains) == 0 {
		fmt.Printf("No chains leading to the package exist!")
	}

	if direct {
		directDep := []string{}
		added := map[string]bool{}
		for _, c := range chains {
			if _, ok := added[c.Pkg]; !ok {
				added[c.Pkg] = true
				directDep = append(directDep, c.Pkg)
			}
		}
		for i, p := range directDep {
			fmt.Printf("%d) %v\n", i+1, p)
		}
	} else {
		for i, c := range chains {
			fmt.Printf("%d) %v\n", i+1, c.String())
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&excludePkg, "exclude", "e", []string{}, "exclude packages that contain the specified string")
	rootCmd.PersistentFlags().BoolVarP(&direct, "direct", "d", false, "only returns unique direct dependencies that use the specified package")
}
