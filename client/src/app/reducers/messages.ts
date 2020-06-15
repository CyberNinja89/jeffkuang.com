import {
    IMessagesState,
    MessageActionTypes,
    ADD_MESSAGE,
    REMOVE_MESSAGE
} from 'app/types/messages'

const initialState: IMessagesState = {
    messages: []
}

export function messageReducer(
    state = initialState,
    action: MessageActionTypes
): IMessagesState {
    switch (action.type) {
        case ADD_MESSAGE:
            return {
                messages: [...state.messages, action.payload]
            }
        case REMOVE_MESSAGE:
            return { messages: state.messages.filter(m => m.id !== action.payload) }
        default:
            return state
    }
}