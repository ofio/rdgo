package rdgo

const QueryApprovalRequest string = `query approval_request($whereCond: approval_request_bool_exp) {
  approval_request(where: $whereCond) {
		id
    uuid
    object_uuid
    cover_page
    instance {
      id
      business {
        id
        name
      }
    }
    message
    status
    external_id
    service_name
    attachment_rank
    approval_request_attachments {
      contract_attachment {
        id
        name
        uuid
        generation
        read_secret
        mime_type
      }
      approval_request_attachment_approvers(order_by: {sequence: asc}) {
        sequence
        is_signer
        approver {
          id
          name
          email
        }
      }
    }
    creator {
      id
      instance_id
      user_preference {
        docusign_refresh_token
        docusign_user_info
        adobe_sign_refresh_token
        adobe_sign_api
      }
    }
    contract {
      id
      currency_code
      note
      annualized_value
      funding_department {
        id
        name
      }
      increase_percent
      renewal_type
      renegotiation_alert_date
      renewal_notification_days
      payment_terms
      effective_date
      end_date
      owner {
        name
      }
      primary_contact {
        name
      }
      managing_department {
        id
        name
      }
      contract_discount_terms {
        discount_days
        discount_percentage
        id
      }
      total_value
      contract_status {
        id
      }
      contract_commodities {
        id
      }
      contract_attachments {
        id
        name
        uuid
        is_deleted
      }
      business {
        id
        name
        phone
      }
      board_item_contracts {
        board_item {
          board {
            board_def
            id
          }
          data
          id
        }
      }
      name
      uuid
      id
    }
    created_by
    instance_id
  }
}
`
