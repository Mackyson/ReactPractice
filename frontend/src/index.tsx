import React from "react"
import ReactDOM from "react-dom"
import Header from "./components/Header.tsx"
import TaskAddingForm from "./components/TaskAddingForm.tsx"
import TaskList from "./components/TaskList.tsx"

ReactDOM.render(
	<div>
		<Header />
		<TaskAddingForm />
		<TaskList />
	</div>,
	document.querySelector("#root")
)
