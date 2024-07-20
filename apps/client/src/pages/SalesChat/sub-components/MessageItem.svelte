<script>
    import { ChatRoomMessage } from '@models/Chat';
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * @type {import('@models/Chat').ChatRoomMessage}
        */
        export let the_message;
        console.log("the_message:", the_message);

        /**
         * Whether the sender is currently typing the message
         * @type {boolean}
         */
        export let is_typing = false;

        /**
         * Time formatter for the message send time
         * @type {Intl.DateTimeFormat}
         */
        const time_formatter = new Intl.DateTimeFormat("en-US", {
            hour: "numeric",
            minute: "numeric",
            hour12: true
        });
            
    
    /*=====  End of Properties  ======*/
    
    
</script>

<li class="scc-message-item" class:recieved-message={is_typing || the_message.Author !== ChatRoomMessage.AUTHOR_USER}>
    <h6 class="sender-name">
        {#if the_message.Author === ChatRoomMessage.AUTHOR_USER}
            You
        {:else}
            {the_message.Author}
        {/if}
    </h6>
    <p class="message-content" class:typing-animation={is_typing}>
        {#if is_typing}
            <span>typing</span>
            <span style:animation-delay=".2s" class="dot"></span>
            <span style:animation-delay=".4s" class="dot"></span>
            <span style:animation-delay=".8s" class="dot"></span>
        {:else}
            {the_message.Content}
        {/if}
    </p>
    <span class="message-metadata">
        <span class="message-send-time">
            {#if is_typing}
                ---
            {:else}
                {time_formatter.format(new Date())}
            {/if}
        </span>
    </span>
</li>

<style>
    .scc-message-item {
        display: flex;
        width: max-content;
        max-width: 95cqw;
        background: var(--theme-color);
        margin: 0 0 0 auto;
        flex-direction: column;
        row-gap: var(--spacing-1);
        padding: var(--spacing-2);
        border-radius: var(--border-radius);
        
        &.recieved-message {
            background: hsl(from var(--theme-hero-background) h s calc(l * 0.5) / 0.3);
            color: var(--theme-color-light);
            margin: 0;
        }
    }

    .recieved-message .sender-name {
        color: var(--accent-orange);
    }

    @keyframes DotLoading {
        20% {
            transform: translateY(-20%);
        }
        50% {
            transform: translateY(0%);
        }
        100% {
            transform: translateY(60%);
        }
    }

    .typing-animation {
        display: flex;
        align-items: flex-end;
        column-gap: calc(var(--spacing-1) * 0.4);

        & > :first-child {
            line-height: 1;
        }

        & > * {
            display: block;
        }
    }

    .message-content .dot {
        width: .4ch;
        height: .4ch;
        border-radius: 50%;
        background: currentColor;
        animation-name: DotLoading;
        animation-duration: 1100ms;
        animation-iteration-count: infinite;
        animation-timing-function: cubic-bezier(0.55, 0.09, 0.68, 0.53);
    }

    .sender-name {
        font-size: var(--font-size-p-small);
        font-weight: bold;
        letter-spacing: 0.08em;
        text-transform: capitalize;
    }

    .message-metadata {
        color: var(--theme-color-light-active);
        line-height: 1.7;
        font-size: var(--font-size-fineprint);
    }


</style>