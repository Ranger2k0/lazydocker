package i18n

func dutchSet() TranslationSet {
	return TranslationSet{
		PruningStatus:              "vernietigen",
		RemovingStatus:             "verwijderen",
		RestartingStatus:           "herstarten",
		StoppingStatus:             "stoppen",
		RunningCustomCommandStatus: "Aangepast commando draaien",

		RunningSubprocess:                          "subprocess draaien",
		NoViewMachingNewLineFocusedSwitchStatement: "No view matching newLineFocused switch statement",

		ErrorOccurred:                     "Er is iets fout gegaan! Zou je hier een issue aan willen maken: https://github.com/jesseduffield/lazydocker/issues",
		ConnectionFailed:                  "connectie naar de docker client mislukt. Het zou kunnen dat je de docker client moet herstarten",
		UnattachableContainerError:        "Container heeft geen ondersteuning voor vastmaken. Je zou de service met het '-it' argument kunnen draaien of stop dit in je `stdin_open: true, tty: true` docker-compose.yml",
		CannotAttachStoppedContainerError: "Je kan niet een vastgemaakte container stoppen, je moet het eerst starten (dit kan je doen met de 'r' toets) (ja ik ben te leu om dat voor je te doen automatisch)",
		CannotAccessDockerSocketError:     "Kan de docker socket niet bereiken: unix:///var/run/docker.sock\nDraai lazydocker als root of lees https://docs.docker.com/install/linux/linux-postinstall/",

		Donate:  "Doneer",
		Confirm: "Bevestigen",

		Return:             "terug",
		FocusMain:          "focus hooft panneel",
		Navigate:           "navigeer",
		Execute:            "voer uit",
		Close:              "sluit",
		Menu:               "menu",
		Scroll:             "scroll",
		OpenConfig:         "open de lazydocker configuratie",
		EditConfig:         "verander de lazydocker configuratie",
		Cancel:             "annuleren",
		Remove:             "verwijder",
		ForceRemove:        "geforceerd verwijderen",
		RemoveWithVolumes:  "verwijder met volumes",
		RemoveService:      "verwijder containers",
		Stop:               "stop",
		Restart:            "herstart",
		Rebuild:            "herbouw",
		Recreate:           "hercreëer",
		PreviousContext:    "vorige tab",
		NextContext:        "volgende tab",
		Attach:             "verbinden",
		ViewLogs:           "bekijk logs",
		RemoveImage:        "verwijder image",
		RemoveVolume:       "verwijder volume",
		RemoveWithoutPrune: "verwijder zonder de ongelabeld ouders te verwijderen",
		PruneContainers:    "vernietig bestaande containers",
		PruneVolumes:       "vernietig ongebruikte volumes",
		PruneImages:        "vernietig ongebruikte images",
		ViewRestartOptions: "beijk herstart opties",
		RunCustomCommand:   "draai een vooraf bedacht aangepaste opdracht",

		AnonymousReportingTitle:  "Help mee met lazydocker beter maken",
		AnonymousReportingPrompt: "Zou je anonieme data rapportage willen aanzetten om lazygit beter te kunnen maken? (enter/esc)",

		GlobalTitle:               "Globaal",
		MainTitle:                 "Hooft",
		ProjectTitle:              "Project",
		ServicesTitle:             "Diensten",
		ContainersTitle:           "Containers",
		StandaloneContainersTitle: "Alleenstaande Containers",
		ImagesTitle:               "Images",
		VolumesTitle:              "Volumes",
		CustomCommandTitle:        "Aangepast commando:",
		ErrorTitle:                "Fout",
		LogsTitle:                 "Logs",
		ConfigTitle:               "Config",
		DockerComposeConfigTitle:  "Docker-Compose Configuratie",
		TopTitle:                  "Top",
		StatsTitle:                "Stats",
		CreditsTitle:              "Over",
		ContainerConfigTitle:      "Container Configuratie",

		NoContainers: "Geen containers",
		NoContainer:  "Geen container",
		NoImages:     "Geen images",
		NoVolumes:    "Geen volumes",

		ConfirmQuit:                "Weet je zeker dat je weg wil gaan?",
		MustForceToRemoveContainer: "Je kan geen draaiende container verwijderen tenzij je het forceert, Wil je het forceren?",
		NotEnoughSpace:             "Niet genoeg ruimte om de panelen te renderen",
		ConfirmPruneImages:         "Weet je zeker dat je alle niet gebruikte images wil vernietigen?",
		ConfirmPruneContainers:     "Weet je zeker dat je alle niet gestopte containers wil vernietigen?",
		ConfirmPruneVolumes:        "Weet je zeker dat je alle niet gebruikte volumes wil vernietigen?",
		StopService:                "Weet je zeker dat je deze service zijn containers wil stoppen (enter/esc)?",
		StopContainer:              "Weet je zeker dat je deze container wil stoppen?",
		PressEnterToReturn:         "Druk op enter om terug te gaan naar lazydocker (Deze popup kan uit gezet worden door in de config dit neer te zetten `gui.returnImmediately: true`)",
	}
}
