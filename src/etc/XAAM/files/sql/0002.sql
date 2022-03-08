INSERT INTO resources(name, created_at, updated_at)
VALUES 
    ('Credit Card', NOW(), NULL),
    ('Debit Cards', NOW(), NULL),
    ('VA', NOW(), NULL),
    ('Retail Outlet', NOW(), NULL),
    ('Wallets', NOW(), NULL),
    ('Cardless Credit', NOW(), NULL),
    ('Local Disbursement' , NOW(), NULL),
    ('Cross-border remittance', NOW(), NULL),
    ('Escrow', NOW(), NULL);

INSERT INTO resource_actions(name, resource_id, created_at, updated_at)
VALUES 
    ('compliant', 1, NOW(), NULL),
    ('compliant', 2, NOW(), NULL),
    ('compliant', 3, NOW(), NULL),
    ('compliant', 4, NOW(), NULL),
    ('compliant', 5, NOW(), NULL),
    ('compliant', 6, NOW(), NULL),
    ('compliant', 7, NOW(), NULL),
    ('compliant', 8, NOW(), NULL),
    ('compliant', 9, NOW(), NULL)
    ('read', 1, NOW(), NULL);

INSERT INTO target_resource_actions(resource_action_id, target_id, target_type, created_at, updated_at)
VALUES 
    (2, 1, 'legal_entity', NOW(), NULL),
    (3, 1, 'legal_entity', NOW(), NULL),
    (4, 1, 'legal_entity', NOW(), NULL),
    (5, 1, 'legal_entity', NOW(), NULL);

INSERT INTO target_resource_actions(resource_action_id, target_id, target_type, created_at, updated_at)
VALUES 
    (7, 2, 'legal_entity', NOW(), NULL),
    (8, 2, 'legal_entity', NOW(), NULL),
    (9, 2, 'legal_entity', NOW(), NULL);

INSERT INTO target_resource_actions(resource_action_id, target_id, target_type, created_at, updated_at)
VALUES 
    (2, 1, 'industry', NOW(), NULL),
    (3, 1, 'industry', NOW(), NULL),
    (4, 1, 'industry', NOW(), NULL),
    (5, 1, 'industry', NOW(), NULL),
    (7, 1, 'industry', NOW(), NULL),
    (8, 1, 'industry', NOW(), NULL),
    (9, 1, 'industry', NOW(), NULL);

