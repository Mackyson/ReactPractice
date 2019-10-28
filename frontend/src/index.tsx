import React from "react"
import ReactDOM from "react-dom"
import Header from "./components/Header.tsx"
import TaskAddingForm from "./components/TaskAddingForm.tsx"
import TaskList from "./components/TaskList.tsx"
import SignUpForm from "./components/SignUpForm.tsx"

ReactDOM.render(
	<div>
		<SignUpForm />
		<Header />
		<TaskAddingForm />
		<TaskList />
	</div>,
	document.querySelector("#root")
)
