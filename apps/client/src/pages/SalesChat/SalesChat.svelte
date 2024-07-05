<script>
    import { onMount } from "svelte";
    import Chatbox from "./sub-components/Chatbox.svelte";
    import ChatStatusBar from "./sub-components/ChatStatusBar.svelte";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        
        /*----------  Styles  ----------*/
        
            const status_bar_height_percentage = 0.1;

            let chat_box_height;
            let status_bar_height;
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        setSubcomponentsHeights();
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/

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
    {#if chat_box_height != null && status_bar_height != null}
        <ChatStatusBar component_pixel_height={status_bar_height} />
        <Chatbox chat_box_pixel_height={chat_box_height} />
    {/if}
</article>