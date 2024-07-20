import { GetChatCredentialsRequest, GetVerifyChatCredentialsRequest, GetChatRoomRequest, PostChatMessageRequest } from "@libs/Services/http_requests";

export class ChatMessage {
    constructor({ chat_id, message, sender, send_date }) {
        this.chat_id = chat_id;
        this.content = message;
        this.sender = sender;
        this.send_date = send_date;
    }
}

class ChatNODE {
    constructor (message, back, next) {
        this.message = message;
        this.back = back;
        this.next = next;
    }
}

export class ChatQueue {
    
    /**
     * @type {ChatNODE}
    */

    #head;
    /**
     * @type {ChatNODE}
    */
    
    #tail;
    /**
     * @type {ChatNODE}
     */
    #traverser;
    constructor() {
        this.#head = null;
        this.#tail = null;
        this.#traverser = null;
    }

    dequeue = () => {
        let node = this.#head;
        this.#head = this.#head?.back;

        if (this.#head) {
            this.#head.next = null;
        } else {
            this.#tail = null;
        }

        return node?.message;
    }

    enqueue = (message) => {
        let node = new ChatNODE(message, null, this.#tail);

        if (this.#tail) {
            this.#tail.back = node;
        }

        this.#tail = node;
        this.#head = this.#head || node;
        
    }

    isEmpty = () => {
        return this.#head === null;
    }

    
    /**
     * Yields the next message in the queue
     * @date 9/11/2023 - 8:26:16 PM
     * @returns {ChatMessage}
    */
    traverse = () => {
        this.#traverser = this.#traverser?.back || this.#head;
        return this.#traverser?.message;
    }

}


/*=============================================
=            Server Chat Classes            =
=============================================*/

export class ChatRoomMessage {
    static AUTHOR_USER = "user";
    static AUTHOR_ASSISTANT = "assistant";

    /**
     * The Order of the message in the chat room where 0 <= order < N(N is a number defined at the chat service. currently there is no way to know the value of N from the client)
     * @type {number}
     */
    #order;

    /**
     * The role who created the message. System messages shouldn't even be sent back to the client to if the message is a system message, an error should be thrown.
     * @type {"user"|"assistant"|"system"}
     */
    #author;

    /**
     * The content of the message
     * @type {string}
     */
    #content;

    /**
     * The date the message was sent in a string with the format "YYYY-MM-DD HH:MM:SS"
     * @type {string}
     */
    #send_date;

    constructor({ order, author, content, send_date }) {
        this.#order = order;
        this.#author = author;
        this.#content = content;
        this.#send_date = send_date;
    }

    /**
     * The order of the message in the chat room.
     * @type {number}
     */
    get Order() {
        return this.#order;
    }

    /**
     * The role of the author of the message.
     * @type {"user"|"assistant"|"system"}
     */
    get Author() {
        return this.#author;
    }

    /**
     * The content of the message
     * @type {string}
     */
    get Content() {
        return this.#content;
    }

    /**
     * Sets the content of a message but only if the message content is empty.
     */
    set Content(new_content) {
        this.#content = new_content;
    }

    /**
     * The date the message was sent in a string with the format "YYYY-MM-DD HH:MM:SS"
     * @type {string}
     */
    get SendDate() {
        return this.#send_date;
    }
}

export class ChatRoom {
    /**
     * The ID of the chat room
     * @type {string}
     */
    #id;

    /**
     * The messages of the chat room.
     * @type {ChatRoomMessage[]}
     */
    #messages;

    /**
     * The max length of a single user message.
     * @type {number}
     */
    #max_user_message_length;

    constructor({ id, messages, max_user_message_length }) {
        this.#id = id;
        this.#messages = messages.map(message => new ChatRoomMessage(message));
        this.#max_user_message_length = max_user_message_length;
    }

    /**
     * Adds a message to the chat room. 
     * @param {ChatRoomMessage} user_message
     * @returns {Promise<ChatRoomMessage>} - the response from the server
     */
    addMessage = async (user_message) => {
        this.#messages.push(user_message);

        const request = new PostChatMessageRequest(user_message.Content);
        const response = await request.do();
        /** @type {ChatRoomMessage} */
        let response_message = null;

        if (response.Ok) {
            response_message = new ChatRoomMessage(response.data);

            this.#messages.push(response_message);
        }

        return response_message;
    };

    /**
     * returns a blank message. this message is not added to the chat room.
     * @param {is_user_message} boolean
     * @returns {ChatRoomMessage}
     */ 
    createBlankMessage = (is_user_message=true) => {
        let author = is_user_message ? ChatRoomMessage.AUTHOR_USER : ChatRoomMessage.AUTHOR_ASSISTANT;
        let send_date = new Date();

        return new ChatRoomMessage({
            order: this.#messages.length,
            author: author,
            content: "",
            send_date: `${send_date.getFullYear()}-${send_date.getMonth() + 1}-${send_date.getDate()} ${send_date.getHours()}:${send_date.getMinutes()}:${send_date.getSeconds()}`
        });
    }

    /**
     * The ID of the chat room
     * @type {string}
     * @readonly
     */
    get ID() {
        return this.#id;
    }

    /**
     * The public messages of the chat room with a different pointer each time it is called, which is useful for reactivity.
     * @type {ChatRoomMessage[]}
     */
    get Messages() {
        let public_messages = [...this.#messages];

        return public_messages; 
    }



}

/*=====  End of Server Chat Classes
  ======*/





/**
 * Verifies if the authentication credentials for the chat service are set and still valid. these are the credentials needed to interact with the chat service and functionality
 * @returns {Promise<boolean>}
 */
export const verifyChatCredentials = async () => {
    let credentials_valid = false;
    const request = new GetVerifyChatCredentialsRequest();
    const response = await request.do();

    if (response.data != null) {
        credentials_valid = response.data.response;
    }

    return credentials_valid;
}

/**
 * Fetches the chat credentials needed to interact with the chat service. If the credentials were not set, it will return an error. otherwise, it will return null  
 * @returns {Promise<Error>}
 */
export const authenticateChatSession = async () => {
    const request = new GetChatCredentialsRequest();
    const response = await request.do();
    let err = null;

    if (response.data == null || !response.data.response) {
        err = new Error('Failed to authenticate chat session');
    }

    return err;    
}

/**
 * Fetches the a chat room from the chat service. This request depends on the chat credentials cookie being set. this is achieved by calling the authenticateChatSession.
 * The verifyChatCredentials will tell you if you need to call the authenticateChatSession to set new credentials.
 * @returns {Promise<ChatRoom>}
 */
export const getChatRoom = async () => {
    const request = new GetChatRoomRequest();
    const response = await request.do();
    /** @type {ChatRoom} */
    let chat_room = null;

    if (response.Ok) {
        chat_room = new ChatRoom(response.data);
    }
    
    return chat_room;
}