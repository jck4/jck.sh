class Terminal {


    constructor() {
        this.chunks = ["Welcome to jck.sh you have been dropped into a shell. Type 'help' for available commands."]
        this.commands = ['help', 'clear', 'info', 'url']
        this.history = []
        this.newLine(this.chunks[0], "terminalLine")
        this.initListeners()
    }

    /*
        Create paragraph element. 
        append element to terminalOutput
        typeEffect String
    */

    newLine(text, lineType) {

        var line = document.createElement("p")
        var terminalOutput = document.getElementById("terminalOutput")
        terminalOutput.appendChild(line)
        line.classList.add(lineType)
        typeEffect(text, 0, text.length, line)
        var terminal = document.getElementById('terminal')
        terminal.scrollTop = terminal.scrollHeight

        return line
    }

    /*
        Stores last command into history array.
        prints last command out to terminal, with previousCommand styling.
    */
    recordHistory(command) {

        if (typeof command === 'string') {
            this.history.push(command)
            var lastCommand = document.createElement("p")
            var terminalOutput = document.getElementById("terminalOutput")
            lastCommand.innerText = "> " + command
            lastCommand.classList.add('previousCommand')
            terminalOutput.appendChild(lastCommand)
        }
        else {
            console.log("Invalid function call recordHistory")
        }

    }


    /*
        Input Handler for terminal input.
        Checks for valid commands 
        performs command functions and 
        controls output the terminal output accordingly
    */

    inputHandler(event, inputField) {

        if (event.key === "Enter") {
            var input = inputField.value
            inputField.value = ""
            this.recordHistory(input)

            switch (input) {
                case "help":
                    this.newLine("Available Commands: " + this.commands, "terminalLine")
                    break;

                case "clear":
                    var outputNode = document.getElementById("terminalOutput")
                    while (outputNode.lastElementChild)
                        outputNode.removeChild(outputNode.lastElementChild);
                    this.chunks = []
                    break;

                case "info":
                    this.newLine("This website is still under construction. More info to come.", "terminalLine")
                    break;

                case "url":
                    this.newLine('ye',"terminalLine")
                    break;

                case "":
                    this.newLine(input, "terminalLine")
                    break;

                default:
                    this.newLine("command not found: " + input, "errorLine")

            }

        }

    }

    initListeners() {
        const input = document.getElementById("terminalTextInput")
        input.addEventListener('keypress', event => this.inputHandler(event, input))
    }

}

document.addEventListener("DOMContentLoaded", function () {
    term = new Terminal()
})