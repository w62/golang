package main

const manifest = `
{
	mainWindow: {
        width:  15%, // screen_width x 15%
        height: 2:1, // width : height = 2 : 1
        title:  "hello",
        content: {
            type: nuxui.org/nuxui/ui.Text,
            text: "Hello nuxui",
		},
	},
}

`
