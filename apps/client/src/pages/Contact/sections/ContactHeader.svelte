<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout"
    import ContactForm from "../sub-components/contact_form.svelte";
</script>

<section id="smk-cp-header-section">
    <div id="smk-cph-wixarica-underlay">
        <WixaricaIcon 
            opacity={0.2}
        />
    </div>
    <div id="smk-cph-content-wrapper" class="design-content-width" class:adebug={false}>
        <div class="wixarica-icons-large">
            <WixaricaIcon />
        </div>
        <header id="smk-cph-header" role="banner">
            <div class="smk-selvamaaku-logo-wrapper">
                <img src="{layout_images.COMBINATION_MARK.getUrl(0.2)}" alt="selvamaaku combination mark" aria-hidden="true">
            </div>
            <hgroup id="smk-cp-header-headlines">
                <h1 id="smk-cp-headline">
                    We are ready to assist you
                </h1>
                <p id="smk-cp-subheadline">
                    Fill out this form and we will be right back at you.
                </p>
            </hgroup>
        </header>
        <div class="wixarica-icons wixarica-icons-large">
            <WixaricaIcon />
        </div>
    </div>
    <div id="smk-cp-contact-section">
        <div id="contact-section-info" class="design-content-width">
            <div id="smk-cp-contact-information-column">
                <div class="smk-cp-isotype-wrapper">
                    <img src="{layout_images.SELVAMAAKU_LOGOTYPE.getUrl(0.1)}" alt="selvamaaku logotype" aria-hidden="true">
                </div>
                <article id="smk-cp-contact-info" role="contentinfo">                    
                    <dl>
                        <dt><strong>Address:</strong></dt>
                        <dd>Km 4.5 Carretera Sayulita – Punta de Mita Sayulita, Nayarit, México.</dd>
                    </dl>
                    <dl>
                        <dt><strong>Sales and info:</strong></dt>
                        <span>Phone +52 322 383 4788</span>
                    </dl>
                    <dl>
                        <dt><strong>Resident attention:</strong></dt>
                        <span>Phone +52 322 383 4788</span>
                    </dl>
                </article>
            </div>
            <div id="smk-cp-contact-form-column">
                <ContactForm />
            </div>
        </div>
    </div>
</section>

<style>
    #smk-cp-header-section {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        row-gap: var(--spacing-5);
    }

    #smk-cph-wixarica-underlay {
        position: absolute;
        bottom: 0;
        left: 50%;
        width: min(var(--design-content-width), 100%);
        transform: translateX(-50%);
    }

    :global(#smk-cp-header-section > #smk-cph-wixarica-underlay svg) {
        mask: linear-gradient(transparent 5%, black 80%);
    }

    #smk-cph-content-wrapper {
        display: flex;
        align-items: center;
        justify-content: center;
        column-gap: var(--spacing-3);
        container-type: inline-size;
    }

    
    /*=============================================
    =            header section            =
    =============================================*/
    
        header#smk-cph-header {
            display: flex;
            flex-direction: column;
            align-items: center;
            row-gap: var(--spacing-3);
        }

        .smk-selvamaaku-logo-wrapper {
            width: min(242px, 19.90131cqw);
        }

        hgroup#smk-cp-header-headlines {
            display: flex;
            flex-direction: column;
            align-items: center;
            row-gap: var(--spacing-2);
            text-align: center;
        }

        h1#smk-cp-headline {
            font-size: var(--font-size-h2);
        }
    
    /*=====  End of header section  ======*/

    
    /*=============================================
    =            Contact section            =
    =============================================*/
    
        #smk-cp-contact-section {
            width: 100%;
        }

        #contact-section-info {
            display: flex;
            justify-content: space-between;
            container-type: inline-size;
        }

        #smk-cp-contact-information-column {
            width: min(281px, 23.10855cqw);   
            display: flex;
            flex-direction: column;
            row-gap: calc(var(--spacing-3) + var(--spacing-1));
        }
        
        .smk-cp-isotype-wrapper {
            width: 100%;
        }

        article#smk-cp-contact-info {
            display: flex;
            flex-direction: column;
            row-gap: calc(var(--spacing-2) + var(--spacing-1));
        }

        article#smk-cp-contact-info dl {
            display: flex;
            flex-direction: column;
            row-gap: var(--spacing-1);

            & dd {
                max-width: 100%;
            }
        }

        
        /*----------  Form column  ----------*/

        #smk-cp-contact-form-column {
            width: min(592px, 48.68421cqw);
        }
        
        


    
    /*=====  End of Contact section  ======*/
    
    
    
    
</style>