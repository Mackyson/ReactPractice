import React from "react"
import ReactDOM from "react-dom"
import MuiThemeProvider from "@material-ui/core/styles"
import Header from "./components/Header.tsx"
import TaskAddingForm from "./components/TaskAddingForm.tsx"

ReactDOM.render(
	<MuiThemeProvider>
		<Header />
		<TaskAddingForm />
	</MuiThemeProvider>,
	document.querySelector("#root")
)
