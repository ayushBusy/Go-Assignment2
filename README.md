# 🏦 Banking System Backend API

A comprehensive, production-grade RESTful API for managing modern banking operations. Built with Go, featuring modern architecture patterns including microservices-ready design, comprehensive error handling, and full support for complex banking scenarios like joint accounts.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Key Features Details](#key-features-details)
- [Configuration](#configuration)
- [Testing](#testing)

---

## Overview

The Banking System Backend API is a robust, scalable solution designed to handle all core banking operations including:

- **Multi-branch banking infrastructure** with hierarchical organization
- **Customer management** with comprehensive profile support
- **Account management** including support for Current, Savings and Joint accounts
- **Transaction processing** with atomic operations and balance management
- **Loan management** with flexible terms and interest calculations [Note: Keeping the interest field empty, will automatically factor it as 12% which is the default case]
- **Repayment tracking** with detailed audit trails

This system is built to support real-world banking scenarios with emphasis on data integrity, security, and operational reliability. All operations maintain ACID compliance through proper transaction handling and database constraints.

---

## Features

### **Bank & Branch Management**

Organize your banking institution into multiple branches for geographic expansion. Each branch operates autonomously while maintaining central oversight. Features include branch-specific information, manager assignments, and administrative hierarchies.

### **Comprehensive Customer Management**

Manage customer information including personal details, contact information, and relationship tracking. The system maintains complete customer profiles and links them to accounts, loans, and transactions for comprehensive financial tracking.

### **Advanced Account Management**

- **Single Accounts**: Individual ownership with full control
- **Joint Accounts**: Multiple account holders with automatic type conversion and role-based management (Primary, Nominee)
- **Account Types**: Support for Savings, Current, and Joint account types
- **Real-time Balance Tracking**: Atomic operations ensure accurate balance management
- **Interest Management**: Commission and interest rate tracking per account

### **Joint Account System** (Advanced Feature)

Joint account functionality that automatically manages account transitions:

- **Automatic Type Conversion**: Account type automatically converts to "joint" when a second customer is added
- **Role-Based Access**: Primary holder vs. joint holder designations
- **Unified Transaction History**: All account holders see the same transactions
- **Flexible Management**: Add or remove joint holders with automatic account type reversion

### **Transaction Management**

Comprehensive transaction tracking with support for:

- **Deposit Operations**: Add funds to accounts
- **Withdrawal Operations**: Remove funds with balance validation
- **Transaction History**: Complete audit trail for compliance
- **Real-time Balance Updates**: Atomic transactions ensure consistency

### **Loan Management System**

Complete loan lifecycle management:

- **Loan Creation**: Define loan terms, amounts, and interest rates
- **Flexible Terms**: Support for various loan durations and structures
- **Interest Calculation**: Automatic interest rate application
- **Loan Status Tracking**: Monitor loan progress from creation to closure

### 📝 **Repayment Tracking**

Detailed repayment functionality:

- **Scheduled Repayments**: Track when repayments are due
- **Payment Records**: Maintain complete audit of all payments made
- **Amount Tracking**: Flexible repayment amounts within loan terms
- **History Maintenance**: Complete payment history for reconciliation

---

##  Tech Stack

| Component         | Technology | Version |
| ----------------- | ---------- | ------- |
| **Language**      | Go         | 1.23.0  |
| **Web Framework** | Gin Gonic  | 1.11.0  |
| **ORM**           | GORM       | 1.31.1  |
| **Database**      | PostgreSQL | Latest  |
| **Driver**        | pgx        | 5.6.0   |
| **Configuration** | godotenv   | 1.5.1   |


---

## 📁 Project Structure

```
banking_system/
├── main.go                          # Application entry point
├── go.mod                           # Go module definition
├── go.sum                           # Dependency lock file
│
├── config/
│   └── db.go                        # Database configuration & initialization
│
├── models/                          # Data models & entities
│   ├── bank.go                      # Bank entity
│   ├── branch.go                    # Branch entity
│   ├── customer.go                  # Customer entity
│   ├── account.go                   # Account entity + AccountDetail response
│   ├── account_customer.go          # Joint account mapping (with Role)
│   ├── loan.go                      # Loan entity
│   ├── repayment.go                 # Repayment entity
│   └── transaction.go               # Transaction entity
│
├── controllers/                     # Request handlers
│   ├── bank_controller.go           # Bank operations
│   ├── branch_controller.go         # Branch operations
│   ├── customer_controller.go       # Customer operations
│   ├── account_controller.go        # Account operations
│   ├── loan_controller.go           # Loan operations
│   ├── repayment_controller.go      # Repayment operations
│   └── transaction_controller.go    # Transaction operations
│
├── services/                        # Business logic layer
│   ├── bank_service.go              # Bank business logic
│   ├── branch_service.go            # Branch business logic
│   ├── customer_service.go          # Customer business logic
│   ├── account_service.go           # Account logic
│   ├── loan_service.go              # Loan business logic
│   ├── repayment_service.go         # Repayment business logic
│   └── transaction_service.go       # Transaction business logic
│
└── routes/
    └── routes.go                    # API route definitions

```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.23.0 or higher
- PostgreSQL 12 or higher
- Git

### Installation

#### 1. Clone the Repository

```bash
git clone <repository-url>
cd Go-Assignment2
```

#### 2. Install Dependencies

```bash
go mod download
go mod verify
```

#### 3. Configure Database

Create a `.env` file in the project root:

```env
DB_URL=postgresql://username:<your_password>@localhost:5432/postgres
PORT=8080
```

#### 4. Initialize Database

```bash
go run main.go
```

The application will automatically:

- Create database tables
- Run migrations
- Seed initial schema

The API will be available at `http://localhost:8080`
