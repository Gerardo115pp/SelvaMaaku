<script>
    import { chat_messages, awaiting_chat_response, chat_room } from "@pages/SalesChat/stores/chat";
    import MessageItem from "./MessageItem.svelte";
    import { ChatRoomMessage } from "@models/Chat";
    import { onDestroy, onMount } from "svelte";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * The height of the component in pixels
         * @type {number}
        */
        export let messages_box_pixel_height;

        /**
         * A chat message object used to represent a typing message(a message been written). null if we are not awaiting a message from the server.
         * @type {ChatRoomMessage}
         * @default null
        */
        let writtenth_message = null;

        let awaiting_new_message_unsubscriber = () => {};
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        awaiting_new_message_unsubscriber = awaiting_chat_response.subscribe(handleAwaitingNewMessageState);
    });

    onDestroy(() => {
        awaiting_new_message_unsubscriber();
    });
    
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Handles changes to the `awaiting_chat_response` store`. if set to true, it sets writtenth_message to blank message. if set to false is sets writtenth_message to null.
         * @param {boolean} new_awaiting_state
        */
        const handleAwaitingNewMessageState = new_awaiting_state => {
            if (new_awaiting_state) {
                let new_writtenth_value = $chat_room.createBlankMessage(false);
                return setTimeout(() => SetWrittenthValue(new_writtenth_value), 300);
            }

            writtenth_message = null;
        }

        /**
         * Sets the value of writtenth_message. meant to be called from SetTimeout.
         * @param {ChatRoomMessage} new_value
        */
        const SetWrittenthValue = new_value => writtenth_message = new_value;
    
    /*=====  End of Methods  ======*/
    
    
    
</script>
<ul id="sales-chat-messanges-list"
    style:height="{messages_box_pixel_height}px"
>
    {#each $chat_messages as message}
        <MessageItem 
            the_message={message}
        />
    {/each}
    {#if writtenth_message !== null} 
        <MessageItem
            the_message={writtenth_message}
            is_typing
        />
    {/if}
</ul>

<style>
    #sales-chat-messanges-list {
        overflow: auto;
        display: flex;
        container-type: inline-size;
        flex-direction: column;
        gap: var(--spacing-3);
        width: 100cqw;
        height: var(--messages-container-height);
        padding: 0 20cqw;
        scrollbar-width: thin;
        scrollbar-color: var(--theme-color) var(--theme-hero-background);
    }
</style>