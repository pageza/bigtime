# User Registration and Authentication Tasks

1. **Backend model and migration**
   - Create `User` model with ID, email, password hash, timestamps
   - Write migration scripts and unit tests

2. **Registration endpoint**
   - Handler to create users
   - Validate email uniqueness and password strength
   - Return JWT on success
   - Swagger docs and handler tests

3. **Login endpoint**
   - Verify credentials and issue JWT
   - Rate limit login attempts
   - Test invalid credentials and locked account cases

4. **Auth middleware**
   - Middleware to validate JWT and set user context
   - Tests for valid/invalid tokens

5. **Frontend screens**
   - Vue components for sign up and sign in
   - Form validation and API integration
   - Unit tests for component logic

6. **README update**
   - Document running auth server, migrations, and endpoints
