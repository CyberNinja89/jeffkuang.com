import { Provider } from "react-redux";
import { BrowserRouter } from "react-router-dom";
import thunk from "redux-thunk";
import { persistStore, persistReducer } from "redux-persist";
import { PersistGate } from "redux-persist/lib/integration/react";
// import reducers from "app/reducers/index";
import {store} from "app/reducers/index";
import "bootstrap/dist/css/bootstrap.css";

import React from 'react';
import ReactDOM from 'react-dom';
import { Router } from "react-router-dom";
import { createBrowserHistory } from "history";

import './index.css';
import App from 'app/layout/App';
import * as serviceWorker from './serviceWorker';
import ScrollToTop from 'app/layout/ScrollToTop';

export const history = createBrowserHistory();

const persistor = persistStore(store);

ReactDOM.render(
    <React.StrictMode>
      <Provider store={store}>
        <PersistGate /*loading={<LoadingScreen />}*/ persistor={persistor}>
        <Router history={history}>
            <ScrollToTop>
                <App />
            </ScrollToTop>
        </Router>
        </PersistGate>
    </Provider>,
    </React.StrictMode>,
    document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();