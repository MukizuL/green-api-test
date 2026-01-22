document.addEventListener("DOMContentLoaded", () => {
	const idInput = document.getElementById("idInstance");
	const tokenInput = document.getElementById("apiTokenInstance");
	const phoneSendMessage = document.getElementById("phoneSendMessage");
	const phoneSendFileByUrl = document.getElementById("phoneSendFileByUrl");

	// Restore saved values
	idInput.value = localStorage.getItem("idInstance") || "";
	tokenInput.value = localStorage.getItem("apiTokenInstance") || "";
	phoneSendMessage.value = localStorage.getItem("phoneSendMessage") || "";
	phoneSendFileByUrl.value = localStorage.getItem("phoneSendFileByUrl") || "";

	// Save on change
	idInput.addEventListener("blur", () => {
		localStorage.setItem("idInstance", idInput.value);
	});

	tokenInput.addEventListener("blur", () => {
		localStorage.setItem("apiTokenInstance", tokenInput.value);
	});

	phoneSendMessage.addEventListener("blur", () => {
		localStorage.setItem("phoneSendMessage", phoneSendMessage.value);
	});

	phoneSendFileByUrl.addEventListener("blur", () => {
		localStorage.setItem("phoneSendFileByUrl", phoneSendFileByUrl.value);
	});
});

async function getSettings() {
	const idInstance = document.getElementById("idInstance").value;
	const apiToken = document.getElementById("apiTokenInstance").value;

	if (!idInstance || !apiToken) {
		alert("Both IdInstance and ApiTokenInstance are required");
		return;
	}

	const response = await fetch("/api/v1/get-settings", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			id_instance: idInstance,
			api_token_instance: apiToken
		})
	});

	const text = await response.text();
	const output = document.getElementById("answer");

	try {
		const json = JSON.parse(text);
		output.value = JSON.stringify(json, null, 2);
	} catch (e) {
		output.value = text;
	}
}

async function getStateInstance() {
	const idInstance = document.getElementById("idInstance").value;
	const apiToken = document.getElementById("apiTokenInstance").value;

	if (!idInstance || !apiToken) {
		alert("Both IdInstance and ApiTokenInstance are required");
		return;
	}

	const response = await fetch("/api/v1/get-state-instance", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			id_instance: idInstance,
			api_token_instance: apiToken
		})
	});

	const text = await response.text();
	const output = document.getElementById("answer");

	try {
		const json = JSON.parse(text);
		output.value = JSON.stringify(json, null, 2);
	} catch (e) {
		output.value = text;
	}
}

async function sendMessage() {
	const idInstance = document.getElementById("idInstance").value;
	const apiToken = document.getElementById("apiTokenInstance").value;
	const phone = document.getElementById("phoneSendMessage").value;
	const message = document.getElementById("sendMessageText").value;

	if (!idInstance || !apiToken) {
		alert("Both IdInstance and ApiTokenInstance are required");
		return;
	}

	const response = await fetch("/api/v1/send-message", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			id_instance: idInstance,
			api_token_instance: apiToken,
			chat_id: phone,
			message: message
		})
	});

	const text = await response.text();
	const output = document.getElementById("answer");

	try {
		const json = JSON.parse(text);
		output.value = JSON.stringify(json, null, 2);
	} catch (e) {
		output.value = text;
	}
}

async function sendFileByUrl() {
	const idInstance = document.getElementById("idInstance").value;
	const apiToken = document.getElementById("apiTokenInstance").value;
	const phone = document.getElementById("phoneSendFileByUrl").value;
	const fileUrl = document.getElementById("fileUrl").value;

	if (!idInstance || !apiToken) {
		alert("Both IdInstance and ApiTokenInstance are required");
		return;
	}

	const response = await fetch("/api/v1/send-file-by-url", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			id_instance: idInstance,
			api_token_instance: apiToken,
			chat_id: phone,
			file_url: fileUrl
		})
	});

	const text = await response.text();
	const output = document.getElementById("answer");

	try {
		const json = JSON.parse(text);
		output.value = JSON.stringify(json, null, 2);
	} catch (e) {
		output.value = text;
	}
}