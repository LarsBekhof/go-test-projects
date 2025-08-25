const generateRandomString = () => Math.floor(Math.random() * Date.now()).toString(36);

const main = () => {
    const socket = new WebSocket("ws://localhost:9000/ws");
    const chat = document.querySelector("#chat");
    const chatbox = document.querySelector("#chatbox");
    const userId = generateRandomString();

    chatbox.addEventListener("keydown", (event) => {
        if (event.key === 'Enter') {
            event.preventDefault();

            if (!chatbox.value) return;

            socket.send(`${userId};${chatbox.value}`);
        }
    });

    socket.addEventListener("open", (event) => {
        console.log(event);
    });

    socket.addEventListener("message", (event) => {
        const [name, message] = event.data.split(';');

        const wrapper = document.createElement("div");

        const nameEl = document.createElement("subtitle")
        nameEl.innerText = name;

        const messageEl = document.createElement("h4")
        messageEl.innerText = message;

        wrapper.appendChild(nameEl);
        wrapper.appendChild(messageEl);

        chat.appendChild(wrapper);
        chatbox.value = '';
    });
};

window.addEventListener("load", main);
