import nodemailer from 'nodemailer';
import Handlebars from 'handlebars';
import fs from 'fs';
import { browser } from '$app/environment';

/**
 * Sends an email. The recipient, and auth data must be present in process.env, e.g process.env.MAIL_USERNAME, process.env.MAIL_PASSWORD, process.env.MAIL_PORT. all of these must be defined.
 * @param {nodemailer.SendMailOptions} mail_options
 * @returns {boolean} true if the email data was correct. doesn't actually check if the email was sent.
 */
const sendMail = mail_options => {
    if (browser) {
        console.warn("sendMail() is not available in the browser.");
        return;
    }

    if (process.env.MAIL_USERNAME === undefined || process.env.MAIL_PASSWORD === undefined || process.env.MAIL_PORT === undefined) {
        console.warn("sendMail() called but the mail server is not configured.")
        return;
    }
    
    const transporter = nodemailer.createTransport({
        service: 'gmail',
        auth: {
            user: process.env.MAIL_USERNAME,
            pass: process.env.MAIL_PASSWORD
        },
        port: parseInt(process.env.MAIL_PORT),
    });

    let send_success = false;

    try {
        transporter.sendMail(mail_options);
        send_success = true;
    } catch (error) {
        console.error(error);
    }

    return send_success;
}

/**
 * Sends a contact mail.
 * @param {string} name
 * @param {string} email
 * @param {string} phone
 * @param {string} message
 * @returns {boolean} true if the email data was correct. doesn't actually check if the email was sent.
 */
export const sendContactMail = (name, email, phone, message) => {
    const contact_template_path = `${process.env.NODE_SERVER_OD}/email/contact_form_template.html`
    let html_body = "";

    if (fs.existsSync(contact_template_path)) {
        const html_source = fs.readFileSync(contact_template_path, "utf8");
        const template = Handlebars.compile(html_source);
        
        html_body = template({ name, email, phone, message });
    } else {
        console.warn(`Contact form template not found at: ${contact_template_path}`)
    }

    const mail_from = process.env.MAIL_FROM ?? "Libery contact service";

    const mail_options = {
        from: `${process.env.MAIL_FROM} <${process.env.MAIL_USERNAME}>`,
        to: process.env.MAIL_CONTACT_EMAIL,
        subject: `Contact form submission from ${name}`,
        text: `Name: ${name}\nEmail: ${email}\nPhone: ${phone}\nMessage: ${message}`,
        html: html_body,
    }

    return sendMail(mail_options);
}