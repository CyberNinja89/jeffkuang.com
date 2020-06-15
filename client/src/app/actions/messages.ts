import { IMessage, ADD_MESSAGE, REMOVE_MESSAGE, MessageActionTypes } from 'app/types/messages'

// TypeScript infers that this function is returning SendMessageAction
export function addMessage(newMessage: IMessage): MessageActionTypes {
    return {
        type: ADD_MESSAGE,
        payload: newMessage
    }
}

export function removeMessage(id: string): MessageActionTypes {
    return {
        type: REMOVE_MESSAGE,
        meta: {
            id
        }
    }
}