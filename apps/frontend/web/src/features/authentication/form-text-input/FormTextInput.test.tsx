import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';
import { describe, it, expect } from 'vitest';
import { FormTextInput } from './FormTextInput';
import userEvent from '@testing-library/user-event';
import React from 'react';

describe('FormTextInput', () => {
	it('should render plain text input if type is text', () => {
		// Arrange
		render(
			<FormTextInput
				placeholder=""
				type="text"
				handleOnChange={() => {}}
			/>
		);

		// Act
		const textInput = screen.getByRole('textbox');
		const showPasswordIcon = screen.queryByAltText(/show password/i);

		// Assert
		expect(showPasswordIcon).not.toBeInTheDocument();
		expect(textInput).toHaveAttribute('type', 'text');
	});

	it('should render password icon and hides input if type is password', () => {
		// Arrange
		render(
			<FormTextInput
				placeholder="Password"
				type="password"
				handleOnChange={() => {}}
			/>
		);

		// Act
		const passwordInput = screen.getByPlaceholderText(/password/i);
		const showPasswordIcon = screen.getByAltText(/show password/i);

		// Assert
		expect(showPasswordIcon).toBeInTheDocument();
		expect(passwordInput).toHaveAttribute('type', 'password');
	});

	it('should expose password as plain-text if show password icon is toggled', async () => {
		// Arrange
		render(
			<FormTextInput
				placeholder=""
				type="password"
				handleOnChange={() => {}}
			/>
		);

		// Act
        const showPasswordIcon = screen.getByAltText(/show password/i);
        await userEvent.click(showPasswordIcon);
        const passwordInput = screen.getByRole('textbox');

		// Assert
		expect(showPasswordIcon).toBeInTheDocument();
		expect(passwordInput).toHaveAttribute('type', 'text');
	});
});
