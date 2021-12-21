---
layout: page
title: File structure
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# File structure

A Metro 2 credit reporting file consists of a single Header Record, Data Record, and Trailer Record. The Data Record includes a Base Segment followed by several optional appended segments. There can be multiple occurences of an optional segment type as needed, except for L1. The Header Record, Base Segment, and Trailer Record can be in either unpacked (character) or packed format, while the optional segments must be unpacked.

Unpacked format for header, base, and trailer segments is a 426 byte sequence that's fixed or variable blocked. Packed format is a 366 byte sequence that's variable blocked. Unpacked format is heavily preferred for electronic transmission.

The following reporting standards apply to all Metro 2 data:

- File names cannot have special characters or whitespace
- Every alphanumeric field must be left-justified and blank filled
- Every alpha field must be uppercase
- If a descriptive field is not available, it should be blank filled
- If a numeric or monetary field is not available, it should be zero filled
- Monetary fields must be reported in whole dollars, with cents truncated
- If fixed-length records are being reported and a record doesn't require the information for an optional appendage segment, the Segment Identifier (e.g., J1) must still be reported and the remainder of the segment must be blank filled

The chart below describes each segment in detail, and segments must occur in the order listed:

| Record/Segment | Description | Length (Bytes) |
| ----------- | ----------- | ----------- |
| Header Record | Identifies the reporter and reporting period. | 426 or 366 |
| Base Segment | Contains account information, which applies to all consumers associated to the account. Also contains information specific to the primary consumer. For joint accounts where two or more consumers are contractually responsible, the consumer reported in the Base Segment will be considered primary for reporting purposes. | 426 or 366 |
| J1 Segment - Associated Consumer (Same Address) | Contains information specific to an associated consumer who lives at the same address as the consumer reported in the Base Segment. All account information reported in the Base Segment will be applied to this consumer. | 100 |
| J2 Segment - Associated Consumer (Different Address) | Contains information specific to an associated consumer who lives at a different address than the consumer reported in the Base Segment. All account information reported in the Base Segment will be applied to this consumer. | 200 |
| K1 Segment - Original Creditor Name | Contains the name of the original credit grantor, including any partnering affinity name, and the creditor’s classification. Reported by collection agencies, debt buyers, check guarantee companies, student loan guaranty agencies, the U.S. Department of Education and the U.S. Treasury. | 34 |
| K2 Segment - Purchased From/Sold To | Contains the name of the company from which an account was purchased or the name of the company to which an account was sold. | 34 |
| K3 Segment - Mortgage Information | Contains the Fannie Mae or Freddie Mac loan number associated to a mortgage account and/or the Mortgage Identification Number assigned by MERS. | 40 |
| K4 Segment Specialized Payment Information | Contains additional account information on deferred payments or balloon payments. | 30 |
| L1 Segment Account Number / Identification Number Change | Used to report a new Account Number and/or new Identification Number. | 54 |
| N1 Segment Employment | Contains employment information for the primary consumer reported in the Base Segment. | 146 |
| Trailer Record | Contains various totals of information reported on the file. | 426 or 366 |
