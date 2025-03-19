# Business logic

## Account types
- Employer-based accounts
  - Set up through employers, but each user (employee) has their own account
- Retail-based (ISA) account
  - A user can sign up for an account directly
  - That user may or may not also have an employer-based account


## Investing into an ISA
- User can't invest more than X amount per year into their ISA
  - These rules should be configurable
  - The rules around how much they can invest are set in configs for now
- For now, user's can't open more than one of the same account type


## ISA allowances 
- ISA allowances are calculated as the total allowance minus the sum of all investments made in the current tax year
- Trying to invest more than the remaining allowance returns an error
- In this hypothetical project, pension contributions don't affect ISA limits for direct accounts 

