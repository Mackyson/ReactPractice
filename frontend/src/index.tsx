import React from "react"
import ReactDOM from "react-dom"
import Header from "./components/Header.tsx"
// import TaskAddingForm from "./components/TaskAddingForm.tsx"
// import TaskList from "./components/TaskList.tsx"
import Task from "./components/Task.tsx"
import SignUpForm from "./components/SignUpForm.tsx"
import SignInForm from "./components/SignInForm.tsx"

ReactDOM.render(
	<div>
		<SignUpForm />
		<SignInForm />
		<Header />
		<Task />
	</div>,
	document.querySelector("#root")
)
