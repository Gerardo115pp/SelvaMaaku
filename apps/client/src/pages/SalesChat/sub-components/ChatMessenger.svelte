<script>
    import { tick } from "svelte";
    import { onMount } from "svelte";
    import { 
        chat_room, 
        is_available, 
        chat_messages,
        awaiting_chat_response
    } from "../stores/chat";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * The message been currently written.
         * @type {import('@models/Chat').ChatRoomMessage}
        */
        let current_message;

        /**
         * Raw message content.
         * @type {string}
        */
       let new_message_content = "";
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        current_message = $chat_room.createBlankMessage(true);        
    });

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Filters keypress events and adds content to the message.
         * @param {KeyboardEvent} event
         * @returns {void}
        */
        const handleMessageWriting = event => {
            const is_enter = event.key === "Enter";

            if ( is_enter && !event.shiftKey) {
                sendMessage();
            }
        }

        /**
         * Sends the message to the chat room.
         * @returns {void}
        */
        const sendMessage = async () => {
            if (new_message_content === "" || !$is_available) return;

            current_message.Content = new_message_content;
            new_message_content = "";
            chat_messages.set([...$chat_messages, current_message])

            
            awaiting_chat_response.set(true);
            // return; // stop here
            let response_message = await $chat_room.addMessage(current_message);
            awaiting_chat_response.set(false);

            await tick();

            chat_messages.set([...$chat_messages, response_message]);
        }
            
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>
<div id="sales-chat-messenger" 
    class:adebug={false}
>
    <div id="scm-messenger-controllers">
        <div id="scm-mc-message-wrapper">
            <textarea type="text" id="message-container" on:keydown={handleMessageWriting} bind:value={new_message_content} rows="1"/>
            <svg id="message-box-tail" viewBox="0 0 50 100">
                <path d="M0 0L0 101H50Q0 101 0 0Z"/>
            </svg>
        </div>
        <button on:click={sendMessage} id="scm-mc-mw-message-sender-btn" class:adebug={false}>
            <svg id="sender-icon" viewBox="0 0 46.297 39.512">
                <path d="m 1.1363858,2.0227688 c 0,0 -1.41023141,-2.26137766 -0.68077897,-1.95840417 C 7.8595857,3.1395619 45.74394,20.288201 45.817448,20.384757 45.898193,20.490833 2.2081654,38.893969 2.2081654,38.893969 c 0,0 -1.0113312,0.364721 -0.4441781,-0.477308 0.3338618,-0.495671 1.5949147,-4.782622 2.63726,-7.697623 1.7646862,-4.93508 4.9435579,-9.667405 5.0704825,-9.667405 0.2018953,0 21.2172532,-0.859524 21.2172532,-0.859524 0,0 1.111157,-0.249633 -0.04038,-0.424281 C 29.732568,19.628897 12.797504,18.388103 8.7459891,18.092292 8.2680232,18.057395 7.7294976,17.670683 7.5431657,17.228448 Z" />
            </svg>
        </button>
    </div>
    <div id="scm-messaenger-disclaims"></div>
</div>

<style>
    #sales-chat-messenger {
        width: 100%;
        height: var(--messenger-height);
        display: grid;
        container-type: inline-size;
    }
    
    /*=============================================
    =            Controllers            =
    =============================================*/
        
        /*----------  Message-input  ----------*/
        
            #scm-messenger-controllers {
                --messenger-sender-btn-width: min(15cqw, 69px);
                --messenger-controllers-column-gap: var(--spacing-1);

                display: grid;
                grid-template-columns: 1fr var(--messenger-sender-btn-width);
                align-items: center;
                column-gap: var(--messenger-controllers-column-gap);
            }

            #scm-mc-message-wrapper {
                --message-box-tail-width: 10.4px;
                --message-box-surface-color: var(--theme-color-light);
                position: relative;
                width: calc(100% - calc(var(--message-box-tail-width) + var(--messenger-controllers-column-gap)));
                /* box-sizing: content-box; */
                background: var(--message-box-surface-color);
                border-radius: var(--border-radius) var(--border-radius) 0 var(--border-radius);
                padding: var(--spacing-1);

                & input, & textarea {
                    width: 100%;
                    background: transparent;
                    border: none;
                    outline: none;
                    color: var(--theme-color-dark);
                }

                & svg {
                    position: absolute;
                    overflow: visible;
                    width: var(--message-box-tail-width);
                    right: calc(var(--message-box-tail-width) * -1); 
                    bottom: 0;
                }

                & svg path {
                    fill: var(--message-box-surface-color);
                }
            }
        
        /*----------  sender-btn  ----------*/

            button#scm-mc-mw-message-sender-btn {
                display: grid;
                width: var(--messenger-sender-btn-width);
                height: var(--messenger-sender-btn-width);
                border-radius: 50%;
                place-items: center;
                background: var(--accent-orange);

                & svg#sender-icon {
                    width: 90%;
                    /* rotate: -45deg; */
                }

                & svg#sender-icon path {
                    fill: var(--theme-color-light);
                }

            }

        @container (width < 500px) {
            #scm-messenger-controllers {
                --messenger-controllers-column-gap: calc(var(--spacing-1) * 0.5)
            }

            button#scm-mc-mw-message-sender-btn {
                padding: calc(var(--spacing-1) * 1.3);
            }

            button#scm-mc-mw-message-sender-btn svg#sender-icon {
                width: 100%;
            }
        }
        
        
        
    
    /*=====  End of Controllers  ======*/
    
    


    


</style>
