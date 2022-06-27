package cmd

/*var getBedTempCmd = &cobra.Command{
	Use:   "getbedtemp",
	Short: "Get bed temp",
	Long:  "Retrieve bed temperature from OctoPrint server",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := octoSvc.GetBedTemp()

		if err != nil {
			log.Println(err)
			return
		}

		quit := make(chan bool)

		go func() {
			for {
				select {
				case <-quit:
					fmt.Println()
					return
				default:
					t, _ = octoSvc.GetBedTemp()
				}
				fmt.Printf("Actual: %.2f, Target: %.2f, Offset: %.2f\r", t.Actual, t.Target, t.Offset)
				time.Sleep(3 * time.Second)
			}
		}()

		// wait for keyboard input to kill thread
		fmt.Scanf("%s")
		quit <- true
	},
}

func init() {
	RootCmd.AddCommand(getBedTempCmd)
}*/
