import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';
import { describe, it, expect } from 'vitest';
import { SignUpForm } from './SignUpForm';
import React from 'react';

describe('SignUpForm', () => {
    it('renders the sign-up form with all required fields', () => {
        render(<SignUpForm />);

        
    });
});