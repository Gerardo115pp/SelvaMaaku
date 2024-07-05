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
