<script>
    import { messages } from "@pages/SalesChat/stores/chat";
    import MessageItem from "./MessageItem.svelte";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * The height of the component in pixels
         * @type {number}
        */
        export let messages_box_pixel_height;
    
    /*=====  End of Properties  ======*/
    
    
</script>
<ul id="sales-chat-messanges-list"
    style:height="{messages_box_pixel_height}px"
>
    {#each $messages as message}
        <MessageItem 
            the_message={message}
            is_typing={message.content === ""}
        />
    {/each}
</ul>

<style>
    #sales-chat-messanges-list {
        container-type: inline-size;
        width: 100cqw;
        height: var(--messages-container-height);
        padding: 0 20cqw;
    }
</style>