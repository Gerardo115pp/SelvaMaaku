<script>
    import { layout_images } from "@stores/layout";
    import ChatMessenger from "./ChatMessenger.svelte";
    import ChatMessages from "./ChatMessages.svelte";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        
        /*----------  Styles  ----------*/
        export let chat_box_pixel_height;
    
    /*=====  End of Properties  ======*/
    
</script>

<section id="sales-chat-chatbox"
    style:height={chat_box_pixel_height + "px"}
    class:adebug={false}
>
    <div id="scc-chatbox-underlay">
        <img src="{layout_images.MANGO_ISOLOGO.getUrl(0.25)}" aria-hidden="true" alt="selvamaaku isologo" class="mango-seal-underlay-wrapper">
    </div>
    <div id="scc-chatbox-messages-wrapper">
        <ChatMessages />
    </div>
    <div id="chat-messenger-wrapper">
        <ChatMessenger />
    </div>
</section>

<style>
    #sales-chat-chatbox {
        --chat-box-gap: var(--spacing-3);
        --chat-box-vpadding: calc(var(--spacing-4) * 2);
        --messages-container-height: calc(80cqh - calc(var(--chat-box-gap)));
        --messenger-height: 20cqh;

        position: relative;
        display: flex;
        width: 100%;
        container-type: size;
        flex-direction: column;
        justify-content: flex-end;
        align-items: center;
        row-gap: var(--chat-box-gap);
        padding: var(--spacing-4) 0;
    }

    #scc-chatbox-underlay {
        position: absolute;
        display: grid;
        width: 100%;
        height: 100%;
        place-items: center;
        inset: 0;
        z-index: var(--z-index-b-1);

        & img {
            width: 20%;
        }
    }

    #chat-messenger-wrapper {
        width: min(690px, 40cqw);
    }
</style>