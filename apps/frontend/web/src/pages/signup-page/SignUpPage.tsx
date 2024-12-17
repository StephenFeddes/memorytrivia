import React from 'react';
import { SignUpForm } from '../../features/authentication/signup-form/SignUpForm';
import styles from './SignUpPage.module.css';

export const SignUpPage = () => {
    return (
        <div className={styles.page}>
            <SignUpForm />
            <div className={styles.description}>
                <header>
                    <p>
                        Memorize card matches and 
                    </p>
                    <p>
                        answer trivia questions
                    </p>
                    <p>
                        to beat your friends
                    </p>
                </header>
            </div>
        </div>
    )
}