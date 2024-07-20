<script>
    import { onMount } from "svelte";
    import Chatbox from "./sub-components/Chatbox.svelte";
    import ChatStatusBar from "./sub-components/ChatStatusBar.svelte";
    import { authenticateChatSession, verifyChatCredentials, getChatRoom } from "@models/Chat";
    import { problem_message, is_available, chat_room, chat_messages } from "./stores/chat";

    
    /*=============================================
    =            Properties            =
    =============================================*/
        
        /*----------  Styles  ----------*/
        
            const status_bar_height_percentage = 0.1;

            let chat_box_height;
            let status_bar_height;
    
    /*=====  End of Properties  ======*/

    onMount(async () => {
        setSubcomponentsHeights();
        
        let err = await bootstrapChat();
        if (err != null) {
            console.error(err);
        }
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/


        /**
         * Bootstraps the chat session and all other necessary components for the chat to work, these been authentications, credentials and chat information.
         * @returns {Promise<Error|null>} Returns an error if there was a problem in the booting process of the chat.
        */
        const bootstrapChat = async () => {
            let authenticath = await authSession();
            
            if (!authenticath) {
                return new Error("There was a problem authenticating the chat session.");
            }

            let new_chat_room = await getChatRoom();
            if (new_chat_room == null) {
                return new Error("There was a problem getting the chat room.");
            }

            chat_room.set(new_chat_room);
            chat_messages.set(new_chat_room.Messages);

            return null;
        }

        /**
         * Verifies whether the authentication credentials are set and if not, requests them.
         * @returns {Promise<boolean>} Returns `true` if the chat credentials are set, `false` otherwise.
         */
        const authSession = async () => {
            let is_verified = await verifyChatCredentials();
            if (!is_verified) {
                let err = await authenticateChatSession();
                is_verified = err == null;
                if (is_verified) {
                    console.error(err);
                    problem_message.set("There was a problem authenticating the chat session. Please try again later.");
                }
            }
            console.log(`Chat session is ${is_verified ? "" : "not "}authenticated.`);

            is_available.set(is_verified);

            return is_verified;
        }


        /**
         * Sets `chat_box_height` and `status_bar_height` based on the window's height and the `status_bar_height_percentage`.
         */
        const setSubcomponentsHeights = () => {
            const root_element = document.querySelector(":root");
            const root_element_styles = getComputedStyle(root_element);
            
            const nabvar_height = parseInt(root_element_styles.getPropertyValue("--navbar-height"));

            const space_left = window.innerHeight - nabvar_height;

            status_bar_height = space_left * status_bar_height_percentage;
            chat_box_height = space_left - status_bar_height;            
        }
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>

<article id="smk-sales-chat">
    {#if chat_box_height != null && status_bar_height != null && $chat_room != null}
        <ChatStatusBar component_pixel_height={status_bar_height} />
        <Chatbox chat_box_pixel_height={chat_box_height} />
    {/if}
</article>