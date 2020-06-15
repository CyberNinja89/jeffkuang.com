import { combineReducers } from 'redux';

import { reducer as formReducer, FormStateMap } from "redux-form";
import {messageReducer} from "./messages";
import {IMessagesState} from "app/types/messages";
import { persistReducer } from 'redux-persist';
import storage from "redux-persist/lib/storage";
import autoMergeLevel2 from "redux-persist/lib/stateReconciler/autoMergeLevel2";
import { createStore, applyMiddleware, compose } from "redux";
import thunk from "redux-thunk";

export interface IRootState {
    formReducer: FormStateMap;
    messageReducer: IMessagesState;
}

const rootReducer = combineReducers<IRootState>({
    formReducer,
    messageReducer,
});

const rootConfig = {
    key: "root",
    storage: storage,
    stateReconciler: autoMergeLevel2,
};

const composeEnhancers = (window && (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose;

const persisted = persistReducer<IRootState>(rootConfig, rootReducer);

export const store = createStore(persisted, composeEnhancers(applyMiddleware(thunk)));

export type RootReducer = ReturnType<typeof rootReducer>;