# Intro to React + Create React App

The purpose of this repo is to get our team better acquainted with React and to officially set us up with a starter boilerplate for building React applications. Ask qustions at anytime!

> A JavaScript library for building user interfaces.
- [ReactJS](https://reactjs.org)


## Brief

React...
- Was released in May of 2009
  - Jordan Walke
- Is currently developed and maintained by Facebook
- Is a library, not a framework
  - Both are code written by another party
  - Different being *Inverted Control*
  - Libary, you are in control
    - Collection of class definitions we can utilize
  - Framework, the framework controls you
    - The Hollywood Principle: "Don't call us, we'll call you!"
    - Change and write our code according to what is given to us
- Is open source
  - Share, modify, publicly available
- Has a large community following
- The most widely used Javascript tool
  - [State of JS 2018 Survey](https://2018.stateofjs.com)
    - Surveyed 20,000 developers worldwide
    - 28% of those surveyed live in US
- Built with Object Oriented Programming paradigm in mind but not tied to it
- Easy to understand
- Component based theory
- Separation of concerns
- Able to power the *V* part in any *MVC* framework
- Makes you a better Javascript developer


## Toolbelt

- HTML
- CSS
- Javascript, ES2016
- JSX, not required but highly recommended
  - It is Javascript, but represents what the UI should look like
- IDE
  - Visual Studio Code, Atom, etc
- Node
- NPM and NPX, both come preshipped with latest Node
- [React JS](https://reactjs.org/)
  - Tutorials
  - Guides
  - Documentation
- [Create React App](https://facebook.github.io/create-react-app/)
  - Note that Webpack and Babel comes preshipped with Create React App package module


## Create React App

Optional, but very highly recommended to use! A toolchain that lets you easily (and painlessly) scaffold out React applications. Get started without any configuration!

To install run
```
npm install -g create-react-app
```

Create React App...
- Can quickly setup up a working boilerplate React app
- Requires very little configuration
- Gives you live-reloading
  - Make edits, save it and see that the browser already has your changes reflected
  - Focus on development
- Offers a great developer experience
- Running Webpack and Babel under the hood, but you won't even notice!


## Things to Know

**Components** are parts that make up a site/application
  - Header, Footer, form controls

**State** is the current situation(state) of a component
  - The state of a checkbox component could be `isChecked = true/false`

**Props** are data/information shared from a parent to a child component
  - Similar to state but is actually passed down to child components

**Life Cycle Methods** custom functionality we can execute throughout all the various stages of a component's life cycle in React


## Getting Started with Hello World!

```
# Open up your terminal app

# Change directory to the desired location
`cd Projects`

# If you have installed Create React App yet
`npm install -g create-react-app`

# Creating the React application
# Note, you can use npm and yarn to scaffold out the project too
`npx create-react-app react-boilerplate`

# Change directory into the project
`cd react-boilerplate`

# Running the project
`npm start`

```


## Lets Build!

Lets create our first Hello world application. Please feel free to follow along.


## Don't try this at home folks, Eject!

Create React App is like magic, but not really too. Lets discover why. Don't try this at home...or rather on your master branch.
```
npm run eject
```

## Reactized

Popular sites partially/completely written in React JS
- [Instagram](https://www.instagram.com/)
- [Airbnb](https://www.airbnb.com/)
- [Dropbox](https://www.dropbox.com/)
- [Paypal](https://www.paypal.com/)
- [Netflix](https://www.netflix.com/)


# Homework

- Famiarize yourself with the basics of React by reading the [glossary page](https://reactjs.org/docs/glossary.html)
- Take this tutorial from [Tania](https://www.taniarascia.com/getting-started-with-react/), it taught me
  - Starting a project
  - States
  - Props
  - Life Cycle Methods
  - Fetching, binding, manipulating, and passing data
  - Building and deploying a React Application




***



This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br>
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (Webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: https://facebook.github.io/create-react-app/docs/code-splitting

### Analyzing the Bundle Size

This section has moved here: https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size

### Making a Progressive Web App

This section has moved here: https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app

### Advanced Configuration

This section has moved here: https://facebook.github.io/create-react-app/docs/advanced-configuration

### Deployment

This section has moved here: https://facebook.github.io/create-react-app/docs/deployment

### `npm run build` fails to minify



This section has moved here: https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify