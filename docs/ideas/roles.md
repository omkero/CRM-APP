Example For a CRM (Customer Relationship Management) system, you’ll need a structured role system where each user has specific responsibilities. Here’s a detailed role hierarchy with permissions:

1. Super Admin (Owner) Can be only One
- Has Full control over the system
- Can create/edit/delete/upgrade users and roles (with a reason)
- Can configure system settings
- Has Access to the server status and data
- Access to all CRM data (not including sensitive data like passwords etc..)
- Has privlige to ban or suspend any one
- Cannot banned everyone or suspend at once (the system will preventing this)
- if super admin is offline over 15 days the system will choose the Most interactive admin
- can give the super user role to any admin (after this decision his role will downgrade to admin)

2. Admin (Admin)
- Can manage and add new emoloyees (add/remove/edit)
- Can create and assign roles (cannot upgrade employee to admin)
- Can view all customer interactions, reports, tasks
- Can see all the current products
- Can change anyone password (only if necessary like employee forgot current password)
- Can suspend and ban any role except admins and Super Admin
- Has Access to all activities not (Super User)
- Can send warnings to any employee or admin
- Has Access to system analytics

3. Sales Manager (Employee)
Can create task to spcific employee
Can send emails to customers if the system has SMTP Service or just use third pasrty service such as gmail etc ..
Can call customers
Has Access and Can create, update, delete and manage products orders sales reports
Can track and Has Access customer interactions
Has Access to Sales/Products/Customers/Employees/Orders/Reports/Marketing analytics

4. Marketing Manager (Employee)
Manages marketing campaigns
Can generate marketing reports
Can send emails to customers if the system has SMTP Service or just use third pasrty service such as gmail etc ..
Can call customers
Can create some tasks
Can access to Marketing analytics data
Has Access to customer interactions
Has Access to products orders sales


5. Customer Support Agent (Employee)
Handles customer complaints and inquiries
Can access customer tickets and respond
Can escalate issues to managers
Can reply to customers
Can see tickets
Can edit/delete/create tickets
Can read send emails if the system has SMTP Service or company email such as gmail,hotmail etc ..


6. Accountant (Employee)
Handles invoices and payments
Can generate financial reports
Can manage refunds and billing issues

7. Product Manager (Employee)
Can add/edit/remove products
Manages product stock and availability
Oversees pricing and promotions
