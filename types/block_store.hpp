namespace koinos { namespace block_store {

// TODO Is there a better name for this data structure than block_item?
struct block_item
{
   /**
    * The hash of the block.
    */
   multihash                      block_id;

   /**
    * The height of the block.
    */
   block_height_type              block_height;

   /**
    * The block data.  If return_block_blob is false, block_blob will be empty.
    */
   opaque< protocol::block >             block;

   /**
    * The block data.  If return_receipt_blob is false, block_receipt_blob will be empty.
    */
   opaque< protocol::block_receipt >     block_receipt;
};

struct block_record
{
   multihash                      block_id;
   block_height_type              block_height;
   std::vector< multihash >       previous_block_ids;

   protocol::block                       block;
   opaque< protocol::block_receipt >     block_receipt;
};

struct transaction_item
{
   protocol::transaction transaction;
};

struct transaction_record
{
   protocol::transaction transaction;
};

} } // koinos::block_store
