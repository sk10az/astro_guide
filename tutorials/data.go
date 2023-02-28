package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {"Welcome", "", welcomeScreen, true},
		"animations": {"Animations",
			"See how to animate components.",
			makeAnimationScreen,
			true,
		},
		"icons": {"Dictionary of astronomy.",
			"Base information about objects of our solar system.",
			iconScreen,
			true,
		},
		"apptabs": {"AppTabs",
			"A container to help divide up an application into functional areas.",
			makeAppTabsTab,
			true,
		},
		"border": {"Border",
			"A container that positions items around a central content.",
			makeBorderLayout,
			true,
		},
		"box": {"Box",
			"A container arranges items in horizontal or vertical list.",
			makeBoxLayout,
			true,
		},
		"center": {"Center",
			"A container to that centers child elements.",
			makeCenterLayout,
			true,
		},
		"doctabs": {"DocTabs",
			"A container to display a single document from a set of many.",
			makeDocTabsTab,
			true,
		},
		"grid": {"Grid",
			"A container that arranges all items in a grid.",
			makeGridLayout,
			true,
		},
		"split": {"Split",
			"A split container divides the container in two pieces that the user can resize.",
			makeSplitTab,
			true,
		},
		"scroll": {"Scroll",
			"A container that provides scrolling for it's content.",
			makeScrollTab,
			true,
		},
		"accordion": {"Accordion",
			"Expand or collapse content panels.",
			makeAccordionTab,
			true,
		},
		"button": {"Button",
			"Simple widget for user tap handling.",
			makeButtonTab,
			true,
		},
		"card": {"Card",
			"Group content and widgets.",
			makeCardTab,
			true,
		},
		"entry": {"Entry",
			"Different ways to use the entry widget.",
			makeEntryTab,
			true,
		},
		"form": {"Form",
			"Gathering input widgets for data submission.",
			makeFormTab,
			true,
		},
		"input": {"Input",
			"A collection of widgets for user input.",
			makeInputTab,
			true,
		},
		"text": {"Text",
			"Text handling widgets.",
			makeTextTab,
			true,
		},
		"toolbar": {"Toolbar",
			"A row of shortcut icons for common tasks.",
			makeToolbarTab,
			true,
		},
		"progress": {"Progress",
			"Show duration or the need to wait for a task.",
			makeProgressTab,
			true,
		},
		"list": {"List",
			"A vertical arrangement of cached elements with the same styling.",
			makeListTab,
			true,
		},
		"table": {"Table",
			"A two dimensional cached collection of cells.",
			makeTableTab,
			true,
		},
		"tree": {"Tree",
			"A tree based arrangement of cached elements with the same styling.",
			makeTreeTab,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "icons"}, // "animations","widgets",
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)
