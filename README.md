# Assisted Writing Application

## Project Summary
This project aims to produce a prototype for an assisted writing application based on [Vale](https://vale.sh/), an exisiting command line interface with a backend in Go. The prototype is developed using [Wails](https://wails.io/) and [Svelte](https://svelte.dev/repl/hello-world). It aims to accept plain text files (.txt or .md) as user input and a preference for style as a text input. Two style preferences are supported in the application with a focus on the word 'data' as either a [singular](./styles/singular) or [plural](./styles/plural) noun. It only captures simple cases for the verbs 'is/are', 'was/were', 'show/shows', 'suggest/suggests', 'indicate/indicates'. However, more complex NLP tools for parts of speech tagging can be included from Python such as [spaCy with Vale](https://github.com/errata-ai/vale/issues/356). The output is displayed in the application with support for errors such as invalid type of text file. 

The application prototype is succesful during development (wails dev) in linking Vale CLI output with user input and displaying errors when for example an image file is read in or incorrect preference is registered like 3 or 4. Although Vale is not able to integrate with Wails during build (wails build) and does not output correctly in the built application, style preferences are captured with correct warnings using the testing text file when using the Vale CLI directly. Future implemntations would explore other prose linters that do not require command line dependencies.

## Important files

**app.go:** Backend 'brain' of application. Lint function takes in user preference and a user-specified text file to output a processed string of the Vale CLI output.

**frontend/src/App.Svelte:** Frontend specifications. Binding of user input for preference to backend code. Displays output of Vale CLI in the application after clicking button.

**.vale.ini**, **styles/singular/singular.yml:** Configuration for singular noun preference and tokens specified

**_vale.ini**, **styles/plural/plural.yml:** Configuration for plural noun preference and tokens specified

**test.txt:** Test file including both 'Data are here' and 'Data is here' with an error produced depending on user preference.

## Installation and Running

Download or git clone this project onto local machine into folder on local machine. Note to substitute brew with your package installer to get Vale. The test file is used to ensure correct installation of vale.
```
git clone https://github.com/asaraog/msds431week7.git
brew install vale
vale sync
vale test.txt --config=.vale.ini
vale test.txt --config=_vale.ini
```

Install Wails. Input 1 or 2 to indicate singular or plural respectively. If correctly input, clicking the 'Lint it' button will prompt a file dialog to open. Select the appropriate text file (test.txt). Output should be identical to using the CLI.
```
xcode-select --install
go install github.com/wailsapp/wails/v2/cmd/wails@latest
wails dev
```

Run application directly.
```
cd ./build/bin/
open Week8.app

```