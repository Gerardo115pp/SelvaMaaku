<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
</script>

<section id="smk-quote-banner-section">
    <div id="smk-qbs-content-wrapper" class="design-content-width">
        <div id="smk-qbs-accent-icon-wrapper" class="wixarica-icons">
            <WixaricaIcon with_accent/>
        </div>
        <blockquote id="smk-qbs-quote-information">
            <h2 id="smk-qbs-qi-quote">
                Is there a reward greater than life?
            </h2>
            <footer>
                <cite id="smk-qbs-quote-source">
                    From the book Life of Pi by Yann Martel
                </cite>
            </footer>
        </blockquote>
    </div>
</section>

<style>
    #smk-quote-banner-section {
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-qbs-content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }

    section#smk-quote-banner-section .wixarica-icons {
        width: calc(var(--wixarica-icons-size) * 1.25);
    }

    blockquote#smk-qbs-quote-information {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-3);
        align-items: center;
        text-align: center;
    }
    
    blockquote#smk-qbs-quote-information h2 {
        font-size: var(--font-size-h3);

        &::before {
            content: "“";
        }

        &::after {
            content: "”";
        }
    }

    blockquote#smk-qbs-quote-information footer cite::before {
        content: "— ";
    }
</style>