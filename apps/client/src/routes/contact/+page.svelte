<script>
    import ContactHeader from "@pages/Contact/sections/ContactHeader.svelte";
    import QuoteBanner from "@pages/Home/sections/QuoteBanner.svelte";
    import { onMount } from "svelte";
    import { pushState, goto } from "$app/navigation";
    import { site_language, supported_languages } from "@stores/site_content";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * The response from the server if the form is submitted. it will be undefined if the form has not been submitted.
         * @type {import('./$types').ActionData|undefined}
         */
        export let form;


        /**
         * The time to wait before redirecting the user to the home page when the form is submitted.
         * @type {number}
         */
        const redirect_timeout = 2000;
    
    /*=====  End of Properties  ======*/
    
    onMount(() => {
        if (form?.success ?? false) {
            prepareRedirect();
        }
    });

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Sets a timeout to redirect the user to the home page
         */
        const prepareRedirect = () => {
            setTimeout(() => {
                goto("/");
            }, redirect_timeout);
        }
    
    /*=====  End of Methods  ======*/

</script>

<main id="smk-contact-page">
    {#if !(form?.success ?? false)}
        <ContactHeader />
        <QuoteBanner />
    {:else}
        <div id="smk-contact-page-success">
            <h1>
                {#if $site_language !== supported_languages.SPANISH}
                    Thank you for reaching out!
                {:else}
                    ¡Gracias por contactarnos!
                {/if}
            </h1>
            <p>
                {#if $site_language !== supported_languages.SPANISH}
                    We will get back to you as soon as possible.
                {:else}
                    Nos pondremos en contacto contigo lo antes posible.
                {/if}
            </p>
        </div>
    {/if}
</main>

<style>
    #smk-contact-page-success {
        display: flex;
        height: calc(100dvh - var(--navbar-height));
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: var(--spacing-3);
        text-align: center;
    }

    #smk-contact-page-success h1 {
        font-size: var(--font-size-h2);
    }

    @media only screen and (max-width: 1200px) {
        #smk-contact-page-success {
            padding: var(--spacing-1);
        }
        
        #smk-contact-page-success h1 {
            font-size: var(--font-size-h1);
        }

        #smk-contact-page-success p {
            font-size: var(--font-size-p-small);
            padding: 0 var(--spacing-1);
        }
    }

    @media only screen and (max-width: 688px) {
        #smk-contact-page-success h1 {
            font-size: calc(var(--font-size-h1) * 0.8);
        }
    }
</style>