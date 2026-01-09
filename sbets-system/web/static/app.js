class SBETSApp {
    constructor() {
        this.init();
    }

    init() {
        this.loadBudget();
        this.loadExpenses();
        this.setupEventListeners();
    }

    setupEventListeners() {
        document.getElementById('expenseForm').addEventListener('submit', (e) => {
            e.preventDefault();
            this.addExpense();
        });
    }

    async loadBudget() {
        try {
            const response = await fetch('/api/budget');
            const data = await response.json();
            
            document.getElementById('budget').innerHTML = `
                <h3>📊 Budget Summary</h3>
                <div class="budget-stats">
                    <div>
                        <div class="total-amount">$${data.totalExpenses.toFixed(2)}</div>
                        <div class="expense-count">${data.expenseCount} expenses</div>
                    </div>
                    <div style="font-size: 2rem;">💵</div>
                </div>
            `;
        } catch (error) {
            console.error('Failed to load budget:', error);
            document.getElementById('budget').innerHTML = `
                <h3>📊 Budget Summary</h3>
                <div class="loading">Failed to load budget</div>
            `;
        }
    }

    async loadExpenses() {
        try {
            const response = await fetch('/api/expenses');
            const expenses = await response.json();
            
            if (!expenses || expenses.length === 0) {
                document.getElementById('expenses').innerHTML = `
                    <div class="no-expenses">
                        <div style="font-size: 3rem; margin-bottom: 10px;">📝</div>
                        <p>No expenses yet. Add your first expense above!</p>
                    </div>
                `;
                return;
            }

            const expensesHTML = expenses.map(expense => `
                <div class="expense-item">
                    <div class="expense-header">
                        <div class="expense-description">${expense.description}</div>
                        <div class="expense-date">${new Date(expense.createdAt).toLocaleDateString()}</div>
                    </div>
                    <div class="expense-amounts">
                        <div class="original-amount">
                            ${expense.amount} ${expense.currency}
                        </div>
                        <div class="converted-amount">
                            → $${expense.convertedAmount.toFixed(2)} USD
                        </div>
                    </div>
                    <button class="delete-btn" onclick="app.deleteExpense(${expense.id})">Delete</button>
                </div>
            `).join('');

            document.getElementById('expenses').innerHTML = expensesHTML;
        } catch (error) {
            console.error('Failed to load expenses:', error);
            document.getElementById('expenses').innerHTML = `
                <div class="loading">Failed to load expenses</div>
            `;
        }
    }

    async addExpense() {
        const amount = parseFloat(document.getElementById('amount').value);
        const currency = document.getElementById('currency').value;
        const description = document.getElementById('description').value;

        if (!amount || !currency || !description) {
            alert('Please fill in all fields');
            return;
        }

        if (amount <= 0) {
            alert('Amount must be greater than zero. Please enter a positive value.');
            return;
        }

        try {
            const response = await fetch('/api/expenses', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    amount: amount,
                    currency: currency,
                    description: description
                })
            });

            if (response.ok) {
                document.getElementById('expenseForm').reset();
                this.loadBudget();
                this.loadExpenses();
                this.showSuccess('Expense added successfully!');
            } else {
                const error = await response.text();
                throw new Error(error);
            }
        } catch (error) {
            console.error('Failed to add expense:', error);
            alert('Failed to add expense: ' + error.message);
        }
    }

    showSuccess(message) {
        // Simple success feedback
        const button = document.querySelector('.btn-primary');
        const originalText = button.textContent;
        button.textContent = '✅ Added!';
        button.style.background = '#48bb78';
        
        setTimeout(() => {
            button.textContent = originalText;
            button.style.background = '';
        }, 2000);
    }

    async deleteExpense(id) {
        if (!confirm('Are you sure you want to delete this expense?')) {
            return;
        }

        try {
            const response = await fetch(`/api/expenses/${id}`, {
                method: 'DELETE'
            });

            if (response.ok) {
                this.loadBudget();
                this.loadExpenses();
            } else {
                const error = await response.text();
                throw new Error(error);
            }
        } catch (error) {
            console.error('Failed to delete expense:', error);
            alert('Failed to delete expense: ' + error.message);
        }
    }
}

// Initialize the app when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    window.app = new SBETSApp();
});