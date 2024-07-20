import { ChatMessage } from "@models/Chat";
import { writable } from "svelte/store";

/**
 * @type {import('svelte/store').Writable<boolean>}
 */
export const is_available = writable(false);

/**
 * @type {import('svelte/store').Writable<string>}
 */
export const problem_message = writable("");

/**
 * Although the messages are also contained in $chat_room, the store is used for reactive updates. and to also filter messages that ar not intended to be displayed.
 * @type {import('svelte/store').Writable<import('@models/Chat').ChatRoomMessage[]>}
 */
export const chat_messages = writable([]);

/**
 * @type {import('svelte/store').Writable<import('@models/Chat').ChatRoom>}
 */
export const chat_room = writable(null);

/** 
 * @type {import('svelte/store').Writable<boolean>}
*/
export const awaiting_chat_response = writable(false);

// Use the following code to initialize the messages store with some data. useful for modifying the chat UI.
// chat_messages.set([
//     new ChatMessage({
//         chat_id: "3b402d3e-09de-42e8-b14c-9d0e6bbbd353",
//         message: "Hello, how can I help you today?",
//         sender: "assistant",
//         send_date: "2023-09-19 20:17:54"  
//     }),
//     new ChatMessage({
//         chat_id: "3b402d3e-09de-42e8-b14c-9d0e6bbbd353",
//         message: "I have a question about your products.",
//         sender: "user",
//         send_date: "2023-09-19 20:18:54"  
//     }),
// ]);

/**
 * @type {import('svelte/store').Writable<boolean>} 
 */
export const is_user_human = writable(false);

/**
 * @type {import('svelte/store').Writable<number>}
 */
export const max_message_length = writable(0);

/**
 * @type {import('svelte/store').Writable<boolean>}
 */
export const is_big_mode_enabled = writable(false);