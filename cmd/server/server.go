package server

/*d
var connectSubCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to OctoPrint server",
	Long:  "Connect to OctoPrint server by server name.  Use `server add` first to create profile.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var status *string
		var err error
		var profile *models.ServerProfile

		useDefault := len(args) == 0

		// load profile based on if user specified to use default or not
		if useDefault {
			profile, err = octoSvc.GetDefaultServerProfile()
		} else {
			profile, err = octoSvc.GetServerProfile(args[0])
		}

		// if any errors loading profile (does not include not finding the profile)
		if err != nil {
			if useDefault {
				fmt.Printf("Error getting default server profile: %s\n", err)
			} else {
				fmt.Printf("Error getting profile for name: '%s'\n", err)
			}
		}

		// if no profile was found
		if profile == nil {
			if useDefault {
				fmt.Println("No default server profile found, create one using `server add`")
			} else {
				fmt.Printf("No server profile found for name: '%s'\n", args[0])
			}
			return
		}

		// connect to server using profile
		status, err = octoSvc.Connect(profile)
		if err != nil {
			fmt.Printf("Error connecting to server: %s\n", err)
			return
		}

		fmt.Printf("Connection status: %s\n", *status)
	},
}
*/
