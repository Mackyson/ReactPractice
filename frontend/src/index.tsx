import React from "react"
import ReactDOM from "react-dom"
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider"
import SignInForm from "./components/SignInForm.tsx"

ReactDOM.render(
	<MuiThemeProvider>
		<SignInForm />
	</MuiThemeProvider>,
	document.querySelector("#root")
)
