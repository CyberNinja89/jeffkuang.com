import React, { Component } from "react";
import { connect } from "react-redux";
import { TransitionGroup, CSSTransition } from "react-transition-group";
import { removeMessage } from "app/actions/messages";
import Toast from "react-bootstrap/Toast";
import ToastHeader from "react-bootstrap/ToastHeader";

// export const messagePosition = {
//     topLeft: { top: 0, left: 0, alignItems: "flex-start", zIndex: 9999 },
//     topCenter: { top: 0, left: "50%", transform: "translate(-50%)", zIndex: 9999 },
//     topRight: { top: 0, right: 0, alignItems: "flex-end", flexWrap: 'wrap-reverse', zIndex: 9999 },
//     bottomLeft: { bottom: 0, left: 0, alignItems: "flex-start", zIndex: 9999 },
//     bottomCenter: { bottom: 0, left: "50%", transform: "translateX(-50%)", zIndex: 9999 },
//     bottomRight: { bottom: 0, right: 0, alignItems: "flex-end", flexWrap: 'wrap-reverse', zIndex: 9999 },
// };

class Messages extends Component {
    componentDidMount() {
        if (this.props.messages && this.props.messages.length > 0) {
            this.props.messages.map((item, index) => {
                this.props.removeMessage(item.id);
            });
        }
    }
    render() {
        let remove = this.props.removeMessage;
        return (
            <div>
                    {this.props.messages.length > 0 &&
                        this.props.messages.map(function (message) {
                            return (
                                <Toast>
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
}

function mapStateToProps(state) {
    return {
        messages: state.messageReducer.messages,
    };
}

const mapDispatchToProps = {
    removeMessage,
};

export default connect(mapStateToProps, mapDispatchToProps)(Messages);
