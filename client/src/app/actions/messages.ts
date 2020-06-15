import { IMessage, ADD_MESSAGE, REMOVE_MESSAGE, MessageActionTypes } from 'app/types/messages'
import { v4 as uuid } from "uuid";

// TypeScript infers that this function is returning SendMessageAction
export function addMessage(newMessage: IMessage): MessageActionTypes {
    return {
        type: ADD_MESSAGE,
        payload: { ...newMessage, id: uuid() }
    }
}

export function removeMessage(id: string): MessageActionTypes {
    return {
        type: REMOVE_MESSAGE,
        payload: id
    }
}