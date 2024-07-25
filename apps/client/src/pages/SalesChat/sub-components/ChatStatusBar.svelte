<script>
    import { is_available } from "../stores/chat";
    import { layout_properties } from "@stores/layout";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        export let component_pixel_height;
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="sales-chat-status-bar"
    style:height={component_pixel_height + "px"}
>
    <h2 id="connection-status" class="chat-status-metric" class:status-failed={!$is_available}>
        {$layout_properties.IS_MOBILE ? "Servicio:" : "Estado de la conexión:"} 
        <span id="connection-status-label" class="status-label">
            {$is_available ? "Conectado" : "Conectando"}
        </span>
    </h2>
</section>

<style>
    #sales-chat-status-bar {
        --status-failed-color: hsl(4, 100%, 58%);
        --status-success-color: hsl(from var(--theme-color) h 100% calc(l * 1.8));

        width: 100%;
        background: var(--theme-color-dark);
        display: flex;
        align-items: center;
        padding: 0 var(--spacing-4);
    }

    .chat-status-metric {
        font-size: var(--font-size-3);

        & span {
            font-family: var(--font-read);
        }
        
        & span.status-label {
            color: var(--status-success-color);
        }

        &.status-failed span.status-label {
            color: var(--status-failed-color);
        }
    }

    @container sales-chat-component (width < 600px) {
        #sales-chat-status-bar {
            padding: 0 var(--spacing-2);
        }

        .chat-status-metric {
            font-size: var(--font-size-2);
        }
    }
</style>