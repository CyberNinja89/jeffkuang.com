import React, { Component } from "react";
import { connect, useSelector, useDispatch } from "react-redux";
// import { TransitionGroup, CSSTransition } from "react-transition-group";
import { removeMessage } from "app/actions/messages";
import Toast from "react-bootstrap/Toast";
import ToastHeader from "react-bootstrap/ToastHeader";
import IMessage from "app/models/message";
import { RootReducer, IRootState } from "app/reducers";
import { IMessagesState, REMOVE_MESSAGE } from "app/types/messages";

// export const messagePosition = {
//     topLeft: { top: 0, left: 0, alignItems: "flex-start", zIndex: 9999 },
//     topCenter: { top: 0, left: "50%", transform: "translate(-50%)", zIndex: 9999 },
//     topRight: { top: 0, right: 0, alignItems: "flex-end", flexWrap: 'wrap-reverse', zIndex: 9999 },
//     bottomLeft: { bottom: 0, left: 0, alignItems: "flex-start", zIndex: 9999 },
//     bottomCenter: { bottom: 0, left: "50%", transform: "translateX(-50%)", zIndex: 9999 },
//     bottomRight: { bottom: 0, right: 0, alignItems: "flex-end", flexWrap: 'wrap-reverse', zIndex: 9999 },
// };
interface Props {
}

interface StateProps {
    messages: IMessage[];
}

const Messages: React.FC = () => {
    const messagesSelector = (state: IRootState) => state.messageReducer.messages;

      const messages = useSelector(messagesSelector);
      const dispatch = useDispatch();

// class Messages extends Component {
    // componentDidMount() {
    //     if (this.props.messages && this.props.messages.length > 0) {
    //         this.props.messages.map((item, index) => {
    //             this.props.removeMessage(item.id);
    //         });
    //     }
        return (
            <div>
                    {messages && messages.length > 0 &&
                        messages.map(function (message) {
                            return (
                                <Toast onClose={() => dispatch({type:REMOVE_MESSAGE, meta: {id: message.id}})}>
                                    <Toast.Header>
                                        <strong className="mr-auto">Notification</strong>
                                    </Toast.Header>
                                    <Toast.Body>{message.text}</Toast.Body>
                                </Toast>
                                // <Fade key={message.id} enter={true} exit={true}>
                                //     <Notification
                                //         type={{ style: message.type, icon: true }}
                                //         closable={true}
                                //         onClose={() => remove(message.id)}
                                //     >
                                //         <span>{message.text}</span>
                                //     </Notification>
                                // </Fade>
                            );
                        })}
            </div>
        );
}

export default Messages;
// function mapStateToProps(state) {
//     return {
//         messages: state.messageReducer.messages,
//     };
// }

// const mapDispatchToProps = {
//     removeMessage,
// };

// export default connect(mapStateToProps, mapDispatchToProps)(Messages);
