export const ADD_MESSAGE = 'ADD_MESSAGE'
export const REMOVE_MESSAGE = 'REMOVE_MESSAGE'

export interface IMessage {
    id?: string
    text: string
    timeout: number
}

export interface IMessagesState {
    messages: IMessage[]
}

interface AddMessageAction {
    type: typeof ADD_MESSAGE
    payload: IMessage
}

interface RemoveMessageAction {
    type: typeof REMOVE_MESSAGE
    payload: string
}

export type MessageActionTypes = AddMessageAction | RemoveMessageAction