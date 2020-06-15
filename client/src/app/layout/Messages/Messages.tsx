import React from "react";
import { useSelector, useDispatch } from "react-redux";
import Toast from "react-bootstrap/Toast";
import { IRootState } from "app/reducers";
import { REMOVE_MESSAGE } from "app/types/messages";

const Messages: React.FC = () => {
    const messagesSelector = (state: IRootState) => state.messageReducer.messages;
    const messages = useSelector(messagesSelector);
    const dispatch = useDispatch();
    return (
        <div>
            {messages && messages.length > 0 &&
                messages.map(function (message, index) {
                    return (
                        <Toast key={index} onClose={() => dispatch({ type: REMOVE_MESSAGE, payload: message.id })}>
                            <Toast.Header>
                                <strong className="mr-auto">Notification</strong>
                            </Toast.Header>
                            <Toast.Body>{message.text}</Toast.Body>
                        </Toast>
                    );
                })}
        </div>
    );
}

export default Messages;
